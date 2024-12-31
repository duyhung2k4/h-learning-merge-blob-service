package streamhandle

import (
	constant "app/internal/constants"
	routerconfig "app/internal/router_config"

	"github.com/gin-gonic/gin"
)

type handleStream struct{}

type HandleStream interface {
	SendBlob(ctx *gin.Context)
}

func NewHandle() HandleStream {
	return &handleStream{}
}

func Register(r *gin.Engine) {
	handle := NewHandle()

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:     constant.GET_HTTP,
		Endpoint:   "stream/blob",
		Middleware: []gin.HandlerFunc{},
		Handle:     handle.SendBlob,
	})
}
