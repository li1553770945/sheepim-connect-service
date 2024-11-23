package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/websocket"
	"github.com/li1553770945/sheepim-connect-service/biz/constant"
	"github.com/li1553770945/sheepim-connect-service/biz/internal/domain"
	"github.com/li1553770945/sheepim-connect-service/kitex_gen/base"
	"github.com/li1553770945/sheepim-connect-service/kitex_gen/message"
)

func (s *MessageService) SendMessage(ctx context.Context, req *message.SendMessageReq) (resp *message.SendMessageResp, err error) {
	klog.CtxInfof(ctx, "请求发送消息到:%s,event:%s,type:%s", req.ClientId, req.Event, req.Type)
	clientId := req.ClientId
	conn, exist := s.ClientConnMap.Get(clientId)
	if !exist || conn == nil {
		return &message.SendMessageResp{
			BaseResp: &base.BaseResp{Code: constant.NotFound, Message: "在本机上不存在的客户端ID"},
		}, nil
	}

	var messageObj = &domain.IMMessageEntity{
		Event: req.Event,
		Type:  req.Type,
		Data:  req.Message,
	}
	messageBytes, err := json.Marshal(messageObj)
	if err != nil {
		klog.CtxErrorf(ctx, "序列化要发送的消息失败：%v", err)
		return &message.SendMessageResp{
			BaseResp: &base.BaseResp{Code: constant.NotFound, Message: fmt.Sprintf("序列化要发送的消息失败：%v", err)},
		}, nil
	}
	err = conn.WriteMessage(websocket.TextMessage, messageBytes)
	if err != nil {
		return nil, err
	}
	return &message.SendMessageResp{
		BaseResp: &base.BaseResp{Code: constant.Success},
	}, nil
}
