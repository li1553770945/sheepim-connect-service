package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/websocket"
	"github.com/li1553770945/sheepim-connect-service/biz/constant"
	"github.com/li1553770945/sheepim-connect-service/biz/middleware"
	"github.com/li1553770945/sheepim-connect-service/biz/model/connect"
	"github.com/li1553770945/sheepim-online-service/kitex_gen/online"
	"strings"
)

var upgrader = websocket.HertzUpgrader{} // use default options

func (s *ConnectService) Connect(ctx context.Context, c *app.RequestContext) *connect.ConnectResp {

	clientId, err := middleware.GetClientIdFromCtx(ctx)
	if err != nil {
		return &connect.ConnectResp{Code: constant.Success, Message: "token认证失败"}
	}

	localIp := utils.LocalIP()
	portIndex := strings.LastIndex(s.Config.ServerConfig.RpcListenAddress, ":")
	if portIndex == -1 {
		hlog.CtxFatalf(ctx, "当前监听路径为%s，无法找到冒号用于分割端口", s.Config.ServerConfig.RpcListenAddress)
		return &connect.ConnectResp{
			Code:    constant.SystemError,
			Message: fmt.Sprintf("当前监听路径为%s，无法找到冒号用于分割端口", s.Config.ServerConfig.RpcListenAddress),
		}
	}
	// 返回从最后一个冒号之后的部分
	port := s.Config.ServerConfig.RpcListenAddress[portIndex+1:]
	endpoint := fmt.Sprintf("%s:%s", localIp, port)
	resp := &connect.ConnectResp{
		Code: constant.Success,
	}
	// TODO:断开连接需要移除客户端
	err = upgrader.Upgrade(c, func(conn *websocket.Conn) {
		hlog.CtxInfof(ctx, "请求升级连接:%s", clientId)
		onlineRpcResp, err := s.OnlineClient.SetClientStatus(ctx, &online.SetClientStatusReq{
			ClientId:       clientId,
			ServerEndpoint: endpoint,
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
			fmt.Println(mt, message, err)
			if err != nil {
				hlog.CtxErrorf(ctx, "读取消息失败:%v", err)
				break
			}
			hlog.CtxInfof(ctx, "收到消息: %s", message)
			err = conn.WriteMessage(mt, message)
			if err != nil {
				hlog.CtxErrorf(ctx, "写入消息失败:%v", err)
				break
			}
		}
		s.ClientConnMap.Remove(clientId)
		onlineRpcResp, err = s.OnlineClient.SetClientStatus(ctx, &online.SetClientStatusReq{
			ClientId:       clientId,
			ServerEndpoint: endpoint,
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
