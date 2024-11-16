package trace

import (
	serverconfig "github.com/cloudwego/hertz/pkg/common/config"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/li1553770945/sheepim-connect-service/biz/infra/config"
)

type TraceStruct struct {
	Provider provider.OtelProvider
	Option   serverconfig.Option
	Config   *hertztracing.Config
}

func InitTrace(config *config.Config) *TraceStruct {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.ServerConfig.ServiceName),
		provider.WithExportEndpoint(config.OpenTelemetryConfig.Endpoint),
		provider.WithInsecure(),
	)
	tracer, cfg := hertztracing.NewServerTracer()
	return &TraceStruct{
		Provider: p,
		Option:   tracer,
		Config:   cfg,
	}

}
