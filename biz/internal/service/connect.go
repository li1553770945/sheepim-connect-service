package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/websocket"
	"github.com/li1553770945/sheepim-connect-service/biz/constant"
	"github.com/li1553770945/sheepim-connect-service/biz/internal/domain"
	"github.com/li1553770945/sheepim-connect-service/biz/middleware"
	"github.com/li1553770945/sheepim-connect-service/biz/model/connect"
	"github.com/li1553770945/sheepim-online-service/kitex_gen/online"
	"github.com/li1553770945/sheepim-push-proxy-service/kitex_gen/push_proxy"
)

var upgrader = websocket.HertzUpgrader{} // use default options

func (s *ConnectService) Connect(ctx context.Context, c *app.RequestContext) *connect.ConnectResp {

	clientId, err := middleware.GetClientIdFromCtx(ctx)
	if err != nil {
		return &connect.ConnectResp{Code: constant.Unauthorized, Message: "token认证失败"}
	}
	roomId := c.Query("roomId")
	if roomId == "" {
		hlog.CtxInfof(ctx, "参数无roomId")
		return &connect.ConnectResp{Code: constant.InvalidInput, Message: "参数无roomId"}
	}

	resp := &connect.ConnectResp{
		Code: constant.Success,
	}
	// TODO:断开连接需要移除客户端
	err = upgrader.Upgrade(c, func(conn *websocket.Conn) {
		hlog.CtxInfof(ctx, "请求升级连接:%s", clientId)
		onlineRpcResp, err := s.OnlineClient.SetClientStatus(ctx, &online.SetClientStatusReq{
			ClientId:       clientId,
			ServerEndpoint: s.Endpoint,
			IsOnline:       true,
		})
		if err != nil {
			hlog.CtxErrorf(ctx, "调用online服务失败:%v", err)
			resp.Code = constant.SystemError
			resp.Message = fmt.Sprintf("调用online服务失败:%v", err)
			return
		}
		if onlineRpcResp.BaseResp.Code != 0 {
			hlog.CtxErrorf(ctx, "调用online服务失败，返回值:%v", onlineRpcResp.BaseResp.Code)
			resp.Code = constant.SystemError
			resp.Message = fmt.Sprintf("调用online服务失败，返回值:%v", onlineRpcResp.BaseResp.Code)
			return
		}

		s.ClientConnMap.Add(clientId, conn)
		hlog.CtxInfof(ctx, "连接成功:%s", clientId)
		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				hlog.CtxErrorf(ctx, "读取消息失败:%v", err)
				break
			}
			hlog.CtxInfof(ctx, "收到来自%s的消息", clientId)

			returnMessage := s.handleMessage(ctx, message, clientId, roomId)

			if returnMessage != nil {
				returnMessageBytes, err := json.Marshal(returnMessage)
				if err != nil {
					klog.CtxErrorf(ctx, "序列化returnMessage失败:%v", err)
					break
				}
				err = conn.WriteMessage(mt, returnMessageBytes)
				if err != nil {
					hlog.CtxErrorf(ctx, "写入消息失败:%v", err)
					break
				}
			}
		}
		s.ClientConnMap.Remove(clientId)
		onlineRpcResp, err = s.OnlineClient.SetClientStatus(ctx, &online.SetClientStatusReq{
			ClientId:       clientId,
			ServerEndpoint: s.Endpoint,
			IsOnline:       false,
		})
		if err != nil {
			hlog.CtxErrorf(ctx, "移除调用online服务失败:%v", err)
			resp.Code = constant.SystemError
			resp.Message = fmt.Sprintf("移除调用online服务失败:%v", err)
		}
		if onlineRpcResp.BaseResp.Code != 0 {
			hlog.CtxErrorf(ctx, "移除调用online服务失败，返回值:%v", onlineRpcResp.BaseResp.Code)
			resp.Code = constant.SystemError
			resp.Message = fmt.Sprintf("移除调用online服务失败，返回值:%v", onlineRpcResp.BaseResp.Code)
		}
	})

	if err != nil {
		hlog.CtxErrorf(ctx, "升级ws连接失败:%v", err)
		return &connect.ConnectResp{Code: constant.SystemError, Message: fmt.Sprintf("升级ws连接失败:%v", err)}
	}
	return nil
}

func (s *ConnectService) handleMessage(ctx context.Context, messageBytes []byte, clientId string, roomId string) *domain.IMMessageEntity {
	var message = &domain.IMMessageEntity{}
	err := json.Unmarshal(messageBytes, message)
	if err != nil {
		return &domain.IMMessageEntity{
			Event: constant.IMError,
			Type:  constant.IMError,
			Data:  fmt.Sprintf("序列化消息失败：%v，消息应包含event,type,data", err),
		}
	}

	if message.Event == constant.IMMessage {
		rpcResp, err := s.PushProxyClient.PushMessage(ctx, &push_proxy.PushMessageReq{
			ClientId: clientId,
			Event:    constant.IMMessage,
			Type:     message.Type,
			RoomId:   roomId,
			Message:  message.Data,
		})
		if err != nil {
			return &domain.IMMessageEntity{
				Event: constant.IMError,
				Type:  constant.IMError,
				Data:  fmt.Sprintf("推送消息到push-proxy失败：%v", err),
			}
		}
		if rpcResp.BaseResp.Code != 0 {
			return &domain.IMMessageEntity{
				Event: constant.IMError,
				Type:  constant.IMError,
				Data:  fmt.Sprintf("推送消息到push-proxy失败：%s", rpcResp.BaseResp.Message),
			}
		}
		return nil
	} else if message.Event == constant.IMPing {
		return &domain.IMMessageEntity{
			Event: constant.IMPong,
			Type:  constant.IMPong,
			Data:  "pong",
		}
	} else {
		return &domain.IMMessageEntity{
			Event: constant.IMError,
			Type:  constant.IMError,
			Data:  "未知的消息event",
		}
	}
}
