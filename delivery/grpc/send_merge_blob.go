package grpchandle

import (
	appcommon "app/cmd/merge-blob/app_common"
	"app/generated/grpc/servicegrpc"
	"errors"
	"io"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (h *grpcHandle) SendMergeBlob(stream grpc.ClientStreamingServer[servicegrpc.SendMergeBlobRequest, servicegrpc.SendMergeBlobResponse]) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("metadata nil")
	}
	uuid := md["uuid"][0]

	socket := appcommon.GetSocket(uuid)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return errors.New("EOF grpc merge-blob-service")
		}

		if err != nil {
			return err
		}

		socket.WriteMessage(websocket.BinaryMessage, req.Blob)
	}
}
