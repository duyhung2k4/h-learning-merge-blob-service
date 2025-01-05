package streamhandle

import (
	constant "app/internal/constants"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func (c *handleStream) StreamM3U8(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	filename := ctx.Param("filename")
	log.Println(filename)
	filepath := fmt.Sprintf("cmd/merge-blob/data/stream/%s/%s", uuid, filename)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		logapp.Logger("get-file-hls", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	ctx.File(filepath)
}
