package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/li1553770945/sheepim-connect-service/biz/infra/config"
	"github.com/li1553770945/sheepim-online-service/kitex_gen/online/onlineservice"
)

func NewOnlineClient(config *config.Config) onlineservice.Client {
	r, err := etcd.NewEtcdResolver(config.EtcdConfig.Endpoint)
	userClient, err := onlineservice.NewClient(
		config.RpcConfig.OnlineServiceName,
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.ServerConfig.ServiceName}),
	)
	if err != nil {
		panic("在线 RPC 客户端启动失败" + err.Error())
	}
	return userClient
}
