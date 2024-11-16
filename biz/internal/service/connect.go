package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/websocket"
	"github.com/li1553770945/sheepim-auth-service/kitex_gen/auth"
	"github.com/li1553770945/sheepim-connect-service/biz/constant"
	"github.com/li1553770945/sheepim-connect-service/biz/model/connect"
	"log"
)

var upgrader = websocket.HertzUpgrader{} // use default options

func (s *ConnectService) Connect(ctx context.Context, c *app.RequestContext) *connect.ConnectResp {

	token, exist := c.Get("token")
	if !exist {
		return &connect.ConnectResp{Code: constant.Success, Message: "token认证失败"}

	}

	tokenStr := token.(string)
	rpcResp, err := s.AuthClient.GetClientId(ctx, &auth.GetClientIdReq{Token: tokenStr})
	if err != nil {
		klog.CtxErrorf(ctx, "调用认证服务失败:%v", err)
		return &connect.ConnectResp{Code: constant.SystemError, Message: fmt.Sprintf("调用认证服务失败:%v", err)}
	}

	if rpcResp.BaseResp.Code != 0 {
		klog.CtxErrorf(ctx, "调用认证服务失败:%s", rpcResp.BaseResp.Message)
		return &connect.ConnectResp{Code: constant.SystemError, Message: fmt.Sprintf("调用认证服务失败:%s", rpcResp.BaseResp.Message)}
	}
	clientId := *rpcResp.ClientId
	// TODO:断开连接需要移除客户端
	err = upgrader.Upgrade(c, func(conn *websocket.Conn) {
		klog.CtxInfof(ctx, "连接成功:%s", clientId)
		s.ClientConnMap.Add(clientId, conn)
		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				klog.CtxErrorf(ctx, "读取失败:%v", err)
				break
			}
			klog.CtxInfof(ctx, "收到消息: %s", message)
			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	})
	if err != nil {
		klog.CtxErrorf(ctx, "升级ws连接失败:%v", err)
		return &connect.ConnectResp{Code: constant.SystemError, Message: fmt.Sprintf("升级ws连接失败:%v", err)}
	}
	return &connect.ConnectResp{Code: constant.Success}
}
