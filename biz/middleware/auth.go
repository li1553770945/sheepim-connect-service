package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	authRpc "github.com/li1553770945/sheepim-auth-service/kitex_gen/auth"
	"github.com/li1553770945/sheepim-auth-service/kitex_gen/auth/authservice"
	"github.com/li1553770945/sheepim-connect-service/biz/constant"
	"strings"
	"sync"
)

func withClientId(ctx context.Context, clientId string) context.Context {
	return context.WithValue(ctx, constant.ClientIdKey, clientId)
}

func GetClientIdFromCtx(ctx context.Context) (string, error) {
	clientId, ok := ctx.Value(constant.ClientIdKey).(string)
	if !ok {
		return "", errors.New("未找到clientId信息")
	}
	return clientId, nil
}

var GlobalAuthMiddleware app.HandlerFunc
var once sync.Once

func GetGlobalGlobalAuthMiddleware() app.HandlerFunc {
	if GlobalAuthMiddleware == nil {
		panic("中间价在使用前未初始化")
	}
	return GlobalAuthMiddleware
}

func InitGlobalAuthMiddleware(authClient authservice.Client) {
	once.Do(func() {
		GlobalAuthMiddleware = AuthMiddleware(authClient)
	})
}

// AuthMiddleware 认证中间件
func AuthMiddleware(authClient authservice.Client) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := string(c.GetHeader("Authorization"))
		if token == "" {
			c.JSON(200, utils.H{"code": constant.Unauthorized, "message": "您还未登陆，请先登录"})
			c.Abort()
			return
		}
		const bearerPrefix = "Bearer "
		if len(token) > len(bearerPrefix) && strings.HasPrefix(token, bearerPrefix) {
			token = token[len(bearerPrefix):]
		}
		req := authRpc.GetClientIdReq{Token: token}
		resp, err := authClient.GetClientId(ctx, &req)
		if err != nil {
			hlog.CtxErrorf(ctx, "认证服务调用失败:%v", err)
			c.JSON(200, utils.H{"code": constant.SystemError, "message": "认证服务调用失败:" + err.Error()})
			c.Abort()
			return
		}
		if resp.BaseResp.Code != 0 {
			hlog.CtxInfof(ctx, "获取客户端ID失败: %v，token 可能已失效", resp.BaseResp.Code)
			c.JSON(200, utils.H{"code": constant.Unauthorized, "message": fmt.Sprintf("获取客户端ID失败: %v，token 可能已失效", resp.BaseResp.Message)})
			c.Abort()
			return
		}

		// 将用户信息添加到上下文中，供后续的处理器使用
		ctx = withClientId(ctx, *resp.ClientId)

		// 调用下一个中间件或处理器
		c.Next(ctx)
	}
}
