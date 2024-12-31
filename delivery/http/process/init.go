package processhandle

import (
	appcommon "app/cmd/merge-blob/app_common"
	"app/generated/grpc/servicegrpc"
	"app/internal/connection"
	constant "app/internal/constants"
	dtoclientservice "app/internal/dto/dto_client_service"
	"app/internal/entity"
	httpresponse "app/pkg/http_response"
	logapp "app/pkg/log"
	"app/pkg/uuidapp"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *processHandle) Init(ctx *gin.Context) {
	var req dtoclientservice.InitProcessRequest
	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		logapp.Logger(constant.TITLE_GET_PAYLOAD, err.Error(), constant.ERROR_LOG)
		httpresponse.BadRequest(ctx, err)
		return
	}

	profileId := ctx.GetUint("profile_id")

	// create chan process
	uuidProcess, err := uuidapp.Create()

	if err != nil {
		logapp.Logger("create-uuid-process", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	newProcessStreamService, err := h.grpcStreamClient.InitProcess(ctx, &servicegrpc.InitProcessRequest{
		Uuid: uuidProcess,
	})
	if err != nil {
		logapp.Logger("init-process-stream", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	newProcessQuantityService, err := h.grpcQuantityClient.InitProcessQuantity(ctx, &servicegrpc.InitProcessQuantityRequest{
		UuidProcess: uuidProcess,
	})
	if err != nil {
		logapp.Logger("init-process-quantity", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	model := entity.ProcessStream{
		ProfileId: profileId,
		IpMergeServer: fmt.Sprintf(
			"%s:%s",
			connection.GetConnect().MergeBlobSevice.Host,
			connection.GetConnect().MergeBlobSevice.Port,
		),
		IpStreamServer: newProcessStreamService.Ip,
		IpQuantity360p: newProcessQuantityService.Ip,
		Uuid:           uuidProcess,
		Status:         entity.PROCESS_PENDING,
	}

	err = appcommon.CreateProcess(uuidProcess)
	if err != nil {
		logapp.Logger("goroutine-process-merge-blob", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	result, err := h.queryProcessStream.Create(model)
	if err != nil {
		logapp.Logger("create-process-merge-blob", err.Error(), constant.ERROR_LOG)
		httpresponse.InternalServerError(ctx, err)
		return
	}

	httpresponse.Success(ctx, result)
}
