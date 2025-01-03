package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/websocket"
	authRpc "github.com/li1553770945/sheepim-auth-service/kitex_gen/auth"
	"github.com/li1553770945/sheepim-connect-service/biz/constant"
	"github.com/li1553770945/sheepim-connect-service/biz/internal/domain"
	"github.com/li1553770945/sheepim-connect-service/biz/model/connect"
	"github.com/li1553770945/sheepim-online-service/kitex_gen/online"
	"github.com/li1553770945/sheepim-push-proxy-service/kitex_gen/push_proxy"
	"github.com/li1553770945/sheepim-room-service/kitex_gen/room"
	"strings"
	"time"
)

var upgrader = websocket.HertzUpgrader{
	CheckOrigin: func(ctx *app.RequestContext) bool {
		return true
	},
} // use default options

func (s *ConnectService) Connect(ctx context.Context, c *app.RequestContext) *connect.ConnectResp {

	roomId := c.Query("roomId")
	if roomId == "" {
		hlog.CtxInfof(ctx, "参数无roomId")
		return &connect.ConnectResp{Code: constant.InvalidInput, Message: "参数无roomId"}
	}

	err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
		hlog.CtxInfof(ctx, "请求升级连接")
		isAuthed := false
		clientId := ""
		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				hlog.CtxErrorf(ctx, "读取消息失败:%v", err)
				err := conn.Close()
				if err != nil {
					hlog.CtxErrorf(ctx, "关闭连接失败:%v", err)
				}
				break
			}

			if isAuthed {
				hlog.CtxInfof(ctx, "收到来自%s的消息", clientId)
			} else {
				hlog.CtxInfof(ctx, "收到来自匿名客户端的消息")
			}

			returnMessage := s.handleEvent(conn, ctx, message, clientId, roomId, isAuthed)

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

				// 根据返回消息进行一些处理
				if returnMessage.Type == constant.IMClose {
					err := conn.Close()
					if err != nil {
						hlog.CtxErrorf(ctx, "关闭连接失败:%v", err)
					}
					break
				}
				if returnMessage.Type == constant.IMAuthResp {
					isAuthed = true
					clientId = returnMessage.Data
				}
			}
		}
		s.ClientConnMap.Remove(clientId)
		onlineRpcResp, err := s.OnlineClient.SetClientStatus(ctx, &online.SetClientStatusReq{
			ClientId:       clientId,
			ServerEndpoint: s.Endpoint,
			IsOnline:       false,
		})
		if err != nil {
			hlog.CtxErrorf(ctx, "移除调用online服务失败:%v", err)
		}
		if onlineRpcResp.BaseResp.Code != 0 {
			hlog.CtxErrorf(ctx, "移除调用online服务失败:%v", onlineRpcResp.BaseResp.Message)
		}
	})

	if err != nil {
		hlog.CtxErrorf(ctx, "升级ws连接失败:%v", err)
		return &connect.ConnectResp{Code: constant.SystemError, Message: fmt.Sprintf("升级ws连接失败:%v", err)}
	}
	return nil
}
func (s *ConnectService) handleEvent(conn *websocket.Conn, ctx context.Context, messageBytes []byte, clientId string, roomId string, isAuthed bool) *domain.IMMessageEntity {
	var message = &domain.IMMessageEntity{}
	err := json.Unmarshal(messageBytes, message)
	if err != nil {
		return &domain.IMMessageEntity{
			Event: constant.IMClose,
			Type:  constant.IMClose,
			Data:  fmt.Sprintf("序列化消息失败：%v，消息应包含event,type,data", err),
		}
	}
	switch message.Event {
	case constant.IMAuthReq:
		{
			token := message.Data
			if token == "" {
				return &domain.IMMessageEntity{
					Event: constant.IMClose,
					Type:  constant.IMClose,
					Data:  fmt.Sprintf("认证失败，token为空"),
				}
			}
			const bearerPrefix = "Bearer "
			if len(token) > len(bearerPrefix) && strings.HasPrefix(token, bearerPrefix) {
				token = token[len(bearerPrefix):]
			}
			req := authRpc.GetClientIdReq{Token: token}
			resp, err := s.AuthClient.GetClientId(ctx, &req)
			if err != nil {
				return &domain.IMMessageEntity{
					Event: constant.IMClose,
					Type:  constant.IMClose,
					Data:  fmt.Sprintf("认证服务调用失败:%s", err.Error()),
				}

			}
			if resp.BaseResp.Code != 0 {
				return &domain.IMMessageEntity{
					Event: constant.IMClose,
					Type:  constant.IMClose,
					Data:  fmt.Sprintf("获取客户端ID失败: %v，token 可能已失效", resp.BaseResp.Message),
				}
			}
			clientId = *resp.ClientId
			roomRpcResp, err := s.RoomClient.CheckIsInRoom(ctx, &room.CheckIsInRoomReq{
				RoomId:   roomId,
				ClientId: clientId,
			})
			if err != nil {
				return &domain.IMMessageEntity{
					Event: constant.IMClose,
					Type:  constant.IMClose,
					Data:  fmt.Sprintf("调用房间rpc服务失败：%v", err),
				}
			}
			if roomRpcResp.BaseResp.Code != 0 {
				return &domain.IMMessageEntity{
					Event: constant.IMClose,
					Type:  constant.IMClose,
					Data:  fmt.Sprintf("调用房间rpc服务失败：%s", roomRpcResp.BaseResp.Message),
				}
			}
			if *roomRpcResp.IsInRoom == false {
				return &domain.IMMessageEntity{
					Event: constant.IMClose,
					Type:  constant.IMClose,
					Data:  fmt.Sprintf("认证失败，您不是该房间的成员"),
				}
			}

			onlineRpcResp, err := s.OnlineClient.SetClientStatus(ctx, &online.SetClientStatusReq{
				ClientId:       clientId,
				ServerEndpoint: s.Endpoint,
				IsOnline:       true,
			})
			if err != nil {
				return &domain.IMMessageEntity{
					Event: constant.IMClose,
					Type:  constant.IMClose,
					Data:  fmt.Sprintf("调用online服务失败:%v", err),
				}
			}
			if onlineRpcResp.BaseResp.Code != 0 {
				return &domain.IMMessageEntity{
					Event: constant.IMClose,
					Type:  constant.IMClose,
					Data:  fmt.Sprintf("调用online服务失败:%v", onlineRpcResp.BaseResp.Message),
				}
			}

			s.ClientConnMap.Add(clientId, conn)
			hlog.CtxInfof(ctx, "连接成功:%s", clientId)
			return &domain.IMMessageEntity{
				Event: constant.IMAuthResp,
				Type:  constant.IMAuthResp,
				Data:  *resp.ClientId,
			}
		}
	case constant.IMMessage:
		{
			if !isAuthed {
				return &domain.IMMessageEntity{
					Event: constant.IMClose,
					Type:  constant.IMClose,
					Data:  fmt.Sprintf("当前客户端未经认证"),
				}
			}
			return s.handleMessage(ctx, message, clientId, roomId)
		}
	case constant.IMPing:
		{
			s.ClientConnMap.UpdateLastPing(clientId)
			onlineRpcResp, err := s.OnlineClient.SetClientStatus(ctx, &online.SetClientStatusReq{
				ClientId:       clientId,
				ServerEndpoint: s.Endpoint,
				IsOnline:       true,
			})
			if err != nil {
				return &domain.IMMessageEntity{
					Event: constant.IMClose,
					Type:  constant.IMClose,
					Data:  fmt.Sprintf("调用online服务失败:%v", err),
				}
			}
			if onlineRpcResp.BaseResp.Code != 0 {
				return &domain.IMMessageEntity{
					Event: constant.IMClose,
					Type:  constant.IMClose,
					Data:  fmt.Sprintf("调用online服务失败:%v", onlineRpcResp.BaseResp.Message),
				}
			}
			return &domain.IMMessageEntity{
				Event: constant.IMPong,
				Type:  constant.IMPong,
				Data:  "im-pong",
			}
		}
	default:
		{
			return &domain.IMMessageEntity{
				Event: constant.IMClose,
				Type:  constant.IMClose,
				Data:  "未知的消息event",
			}
		}
	}
}
func (s *ConnectService) handleMessage(ctx context.Context, message *domain.IMMessageEntity, clientId string, roomId string) *domain.IMMessageEntity {
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
}

func (s *ConnectService) RemoveInactiveClients() {
	for {
		time.Sleep(10 * time.Second) // 每 10 秒检查一次
		klog.Infof("开始进行超时检测")
		s.ClientConnMap.mu.Lock()
		now := time.Now()
		for clientID, lastPingTime := range s.ClientConnMap.lastPingTimeMap {
			if now.Sub(lastPingTime) > constant.HearBearTimeOut {
				// 超时，移除客户端
				err := s.ClientConnMap.store[clientID].Close()
				if err != nil {
					klog.Errorf("%s超时关闭连接失败:%v", clientID, err)
				}
				delete(s.ClientConnMap.store, clientID)
				delete(s.ClientConnMap.lastPingTimeMap, clientID)
				klog.Infof("客户端  %s 超时被移除", clientID)
				rpcResp, err := s.OnlineClient.SetClientStatus(context.TODO(), &online.SetClientStatusReq{
					ClientId:       clientID,
					ServerEndpoint: s.Endpoint,
					IsOnline:       false,
				})
				if err != nil {
					if err != nil {
						klog.Errorf("%s超时关闭连接，调用online服务失败:%v", clientID, err)
					}
				}
				if rpcResp.BaseResp.Code != 0 {
					klog.Errorf("%s超时关闭连接，调用online服务失败:%s", clientID, rpcResp.BaseResp.Message)
				}
			}
		}
		s.ClientConnMap.mu.Unlock()
	}
}
