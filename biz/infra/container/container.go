package container

import (
	"github.com/li1553770945/sheepim-auth-service/kitex_gen/auth/authservice"
	"github.com/li1553770945/sheepim-connect-service/biz/infra/config"
	"github.com/li1553770945/sheepim-connect-service/biz/infra/log"
	"github.com/li1553770945/sheepim-connect-service/biz/infra/trace"
	"github.com/li1553770945/sheepim-connect-service/biz/internal/service"
	"sync"
)

type Container struct {
	TraceStruct    *trace.TraceStruct
	Logger         *log.TraceLogger
	Config         *config.Config
	ConnectService service.IConnectService
	MessageService service.IMessageService
	AuthRpcClient  authservice.Client
}

var APP *Container
var once sync.Once

func GetGlobalContainer() *Container {
	if APP == nil {
		panic("APP在使用前未初始化")
	}
	return APP
}

func InitGlobalContainer(env string) {
	once.Do(func() {
		APP = GetContainer(env)
	})
}

func NewContainer(config *config.Config,
	logger *log.TraceLogger,
	traceStruct *trace.TraceStruct,

	authRpcClient authservice.Client,

	connectService service.IConnectService,
	messageService service.IMessageService,
) *Container {
	return &Container{
		Config:        config,
		Logger:        logger,
		TraceStruct:   traceStruct,
		AuthRpcClient: authRpcClient,

		ConnectService: connectService,
		MessageService: messageService,
	}

}
