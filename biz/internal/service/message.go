package service

import (
	"context"
	"github.com/hertz-contrib/websocket"
	"github.com/li1553770945/sheepim-connect-service/biz/constant"
	"github.com/li1553770945/sheepim-connect-service/kitex_gen/base"
	"github.com/li1553770945/sheepim-connect-service/kitex_gen/message"
)

func (s *MessageService) SendMessage(ctx context.Context, req *message.SendMessageReq) (resp *message.SendMessageResp, err error) {

	clientId := req.ClientId
	conn, exist := s.ClientConnMap.Get(clientId)
	if !exist || conn == nil {
		// TODO:
		return &message.SendMessageResp{
			BaseResp: &base.BaseResp{Code: constant.NotFound, Message: "在本机上不存在的客户端ID"},
		}, nil
	}
	err = conn.WriteMessage(websocket.TextMessage, []byte(req.Message))
	if err != nil {
		return nil, err
	}
	return &message.SendMessageResp{
		BaseResp: &base.BaseResp{Code: constant.Success},
	}, nil
}
