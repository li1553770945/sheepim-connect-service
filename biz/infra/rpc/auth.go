package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/li1553770945/sheepim-auth-service/kitex_gen/auth/authservice"
	"github.com/li1553770945/sheepim-connect-service/biz/infra/config"
)

func NewAuthClient(config *config.Config) authservice.Client {
	r, err := etcd.NewEtcdResolver(config.EtcdConfig.Endpoint)
	userClient, err := authservice.NewClient(
		config.RpcConfig.AuthServiceName,
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.ServerConfig.ServiceName}),
	)
	if err != nil {
		panic("认证 RPC 客户端启动失败" + err.Error())
	}
	return userClient
}
