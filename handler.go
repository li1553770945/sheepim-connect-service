package main

import (
	"context"
	message "github.com/li1553770945/sheepim-connect-service/kitex_gen/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// SendMessage implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) SendMessage(ctx context.Context, req *message.SendMessageReq) (resp *message.SendMessageResp, err error) {
	// TODO: Your code here...
	return
}
