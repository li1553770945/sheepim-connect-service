package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/li1553770945/sheepim-auth-service/kitex_gen/auth/authservice"
	"github.com/li1553770945/sheepim-connect-service/biz/infra/config"
	"github.com/li1553770945/sheepim-connect-service/biz/model/connect"
	"github.com/li1553770945/sheepim-connect-service/kitex_gen/message"
	"github.com/li1553770945/sheepim-online-service/kitex_gen/online/onlineservice"
)

type ConnectService struct {
	ClientConnMap *ClientConnMap
	AuthClient    authservice.Client
	OnlineClient  onlineservice.Client
	Config        *config.Config
}

type IConnectService interface {
	Connect(ctx context.Context, c *app.RequestContext) *connect.ConnectResp
}

func NewConnectService(clientConnMap *ClientConnMap,
	authClient authservice.Client,
	onlineClient onlineservice.Client,
	cfg *config.Config,
) IConnectService {
	return &ConnectService{
		ClientConnMap: clientConnMap,
		AuthClient:    authClient,
		OnlineClient:  onlineClient,
		Config:        cfg,
	}
}

type MessageService struct {
	ClientConnMap *ClientConnMap
}

type IMessageService interface {
	SendMessage(ctx context.Context, req *message.SendMessageReq) (resp *message.SendMessageResp, err error)
}

func NewMessageService(clientConnMap *ClientConnMap) IMessageService {
	return &MessageService{
		ClientConnMap: clientConnMap,
	}
}
