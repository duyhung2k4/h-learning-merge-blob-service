package streamhandle

import (
	appcommon "app/cmd/merge-blob/app_common"
	constant "app/internal/constants"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *handleStream) SendBlob(ctx *gin.Context) {
	uuid := ctx.Query("uuid")

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		httpresponse.InternalServerError(ctx, err)
		logapp.Logger("up-connection", err.Error(), constant.ERROR_LOG)
		return
	}
	defer conn.Close()

	appcommon.CreateSocket(uuid, conn)

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Lỗi khi nhận tin nhắn từ WebSocket: %v", err)
			break
		}

		log.Println(string(data))
	}
}
