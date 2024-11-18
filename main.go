// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"context"
	hertzserver "github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	kitexserver "github.com/cloudwego/kitex/server"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/li1553770945/sheepim-connect-service/biz/infra/container"
	"github.com/li1553770945/sheepim-connect-service/biz/middleware"
	"github.com/li1553770945/sheepim-connect-service/global_middleware"
	"github.com/li1553770945/sheepim-connect-service/kitex_gen/message/messageservice"
	"net"
	"os"
)

func ListenWs() {
	App := container.GetGlobalContainer()
	addr, err := net.ResolveTCPAddr("tcp", App.Config.ServerConfig.WsListenAddress)
	if err != nil {
		panic("设置监听地址出错")
	}

	h := hertzserver.Default(
		App.TraceStruct.Option,
		hertzserver.WithHostPorts(addr.String()),
	)

	h.Use(hertztracing.ServerMiddleware(App.TraceStruct.Config))
	h.Use(global_middleware.TraceIdMiddleware())
	register(h)
	h.Spin()
}
func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	container.InitGlobalContainer(env)
	App := container.GetGlobalContainer()
	middleware.InitGlobalAuthMiddleware(App.AuthRpcClient)

	serviceName := App.Config.ServerConfig.ServiceName

	defer func(p provider.OtelProvider, ctx context.Context) {
		err := p.Shutdown(ctx)
		if err != nil {
			klog.Fatalf("kitexserver stopped with error:%s", err)
		}
	}(App.TraceStruct.Provider, context.Background())

	addr, err := net.ResolveTCPAddr("tcp", App.Config.ServerConfig.RpcListenAddress)
	if err != nil {
		panic("设置监听地址出错")
	}

	r, err := etcd.NewEtcdRegistry(App.Config.EtcdConfig.Endpoint) // r should not be reused.
	if err != nil {
		panic(err)
	}
	svr := messageservice.NewServer(
		new(MessageServiceImpl),
		kitexserver.WithSuite(tracing.NewServerSuite()),
		kitexserver.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
		kitexserver.WithRegistry(r),
		kitexserver.WithServiceAddr(addr),
	)
	go ListenWs()
	if err := svr.Run(); err != nil {
		klog.Fatalf("服务启动失败:", err)
	}

}
