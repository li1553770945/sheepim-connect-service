// Code generated by hertz generator.

package connect

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/li1553770945/sheepim-connect-service/biz/middleware"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _connectMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		middleware.GetGlobalGlobalAuthMiddleware(),
	}
}
