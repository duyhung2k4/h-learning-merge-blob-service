package processhandle

import (
	appcommon "app/cmd/merge-blob-service/app_common"
	constant "app/internal/constants"
	dtoclientservice "app/internal/dto/dto_client_service"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

func (h *processHandle) AddData(ctx *gin.Context) {
	var req dtoclientservice.AddData
	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		httpresponse.BadRequest(ctx, err)
		return
	}

	process := appcommon.GetProcessStream(req.UuidProcess)
	if process == nil {
		httpresponse.InternalServerError(ctx, errors.New("process not found"))
		return
	}
	process <- req.Mess

	httpresponse.Success(ctx, nil)
}
