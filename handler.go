package main

import (
	"context"
	"github.com/li1553770945/sheepim-connect-service/biz/infra/container"
	message "github.com/li1553770945/sheepim-connect-service/kitex_gen/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// SendMessage implements the MessageServiceImpl interface.
func (s MessageServiceImpl) SendMessage(ctx context.Context, req *message.SendMessageReq) (resp *message.SendMessageResp, err error) {
	App := container.GetGlobalContainer()
	resp, err = App.MessageService.SendMessage(ctx, req)
	return
}
