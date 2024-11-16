// Code generated by hertz generator.

package connect

import (
	"context"
	"github.com/li1553770945/sheepim-connect-service/biz/infra/container"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Connect .
// @router /connect [GET]
func Connect(ctx context.Context, c *app.RequestContext) {
	App := container.GetGlobalContainer()
	resp := App.ConnectService.Connect(ctx, c)
	c.JSON(consts.StatusOK, resp)
}
