// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package container

import (
	"github.com/li1553770945/sheepim-connect-service/biz/infra/config"
	"github.com/li1553770945/sheepim-connect-service/biz/infra/log"
	"github.com/li1553770945/sheepim-connect-service/biz/infra/rpc"
	"github.com/li1553770945/sheepim-connect-service/biz/infra/trace"
	"github.com/li1553770945/sheepim-connect-service/biz/internal/service"
)

// Injectors from wire.go:

func GetContainer(env string) *Container {
	configConfig := config.GetConfig(env)
	traceLogger := log.InitLog()
	traceStruct := trace.InitTrace(configConfig)
	clientConnMap := service.NewClientConnMap()
	client := rpc.NewAuthClient(configConfig)
	iConnectService := service.NewConnectService(clientConnMap, client)
	iMessageService := service.NewMessageService(clientConnMap)
	container := NewContainer(configConfig, traceLogger, traceStruct, iConnectService, iMessageService)
	return container
}
