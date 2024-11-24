package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/websocket"
	"github.com/li1553770945/sheepim-auth-service/kitex_gen/auth/authservice"
	"github.com/li1553770945/sheepim-connect-service/biz/infra/config"
	"github.com/li1553770945/sheepim-connect-service/biz/infra/utils"
	"github.com/li1553770945/sheepim-connect-service/biz/internal/domain"
	"github.com/li1553770945/sheepim-connect-service/biz/model/connect"
	"github.com/li1553770945/sheepim-connect-service/kitex_gen/message"
	"github.com/li1553770945/sheepim-online-service/kitex_gen/online/onlineservice"
	"github.com/li1553770945/sheepim-push-proxy-service/kitex_gen/push_proxy/pushproxyservice"
	"github.com/li1553770945/sheepim-room-service/kitex_gen/room/roomservice"
	"strings"
)

type ConnectService struct {
	ClientConnMap   *ClientConnMap
	AuthClient      authservice.Client
	OnlineClient    onlineservice.Client
	PushProxyClient pushproxyservice.Client
	RoomClient      roomservice.Client
	Config          *config.Config
	Endpoint        string
}

type IConnectService interface {
	Connect(ctx context.Context, c *app.RequestContext) *connect.ConnectResp
	handleMessage(ctx context.Context, message *domain.IMMessageEntity, clientId string, roomId string) *domain.IMMessageEntity
	handleEvent(conn *websocket.Conn, ctx context.Context, messageBytes []byte, clientId string, roomId string, isAuthed bool) *domain.IMMessageEntity
}

func NewConnectService(clientConnMap *ClientConnMap,
	authClient authservice.Client,
	onlineClient onlineservice.Client,
	pushProxyClient pushproxyservice.Client,
	roomClient roomservice.Client,
	cfg *config.Config,
) IConnectService {
	localIpList, err := utils.GetLocalIP()

	if err != nil {
		panic(fmt.Sprintf("获取本机ip失败：%v", err))
	}

	localIp := localIpList[len(localIpList)-1]
	portIndex := strings.LastIndex(cfg.ServerConfig.RpcListenAddress, ":")
	if portIndex == -1 {
		panic(fmt.Sprintf("当前监听路径为%s，无法找到冒号用于分割端口", cfg.ServerConfig.RpcListenAddress))

	}
	// 返回从最后一个冒号之后的部分
	port := cfg.ServerConfig.RpcListenAddress[portIndex+1:]
	endpoint := fmt.Sprintf("%s:%s", localIp, port)
	klog.Warnf("当前选择的endpoint为：%s，请确认是否正确！", endpoint)
	return &ConnectService{
		ClientConnMap:   clientConnMap,
		AuthClient:      authClient,
		OnlineClient:    onlineClient,
		PushProxyClient: pushProxyClient,
		RoomClient:      roomClient,
		Config:          cfg,
		Endpoint:        endpoint,
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
