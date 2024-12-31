package processhandle

import (
	appcommon "app/cmd/merge-blob/app_common"
	"app/generated/grpc/servicegrpc"
	"app/internal/connection"
	constant "app/internal/constants"
	"app/internal/entity"
	middlewareapp "app/internal/middleware"
	routerconfig "app/internal/router_config"
	query "app/pkg/query/basic"
	rawquery "app/pkg/query/raw"

	"github.com/gin-gonic/gin"
)

type processHandle struct {
	chanListenAddProcessStream (chan string)
	grpcStreamClient           servicegrpc.StreamServiceClient
	queryProcessStream         query.QueryService[entity.ProcessStream]
	rawQueryProcessStream      rawquery.QueryRawService[entity.ProcessStream]
}

type ProcessHandle interface {
	Init(ctx *gin.Context)
	AddData(ctx *gin.Context)
}

func NewHandle() ProcessHandle {
	return &processHandle{
		chanListenAddProcessStream: appcommon.GetChanListenAddProcessStream(),
		grpcStreamClient:           connection.GetGrpcClientStream(),
		queryProcessStream:         query.Register[entity.ProcessStream](),
		rawQueryProcessStream:      rawquery.Register[entity.ProcessStream](),
	}
}

func Register(r *gin.Engine) {
	handle := NewHandle()

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:   constant.POST_HTTP,
		Endpoint: "process-stream/init",
		Middleware: []gin.HandlerFunc{
			middlewareapp.ValidateToken,
			middlewareapp.GetProfileId,
		},
		Handle: handle.Init,
	})

	routerconfig.AddRouter(r, routerconfig.RouterConfig{
		Method:     constant.POST_HTTP,
		Endpoint:   "process-stream/add-data",
		Middleware: []gin.HandlerFunc{},
		Handle:     handle.AddData,
	})
}
