package grpchandle

import "app/generated/proto/servicegrpc"

type grpcHandle struct {
	servicegrpc.UnimplementedMergeBlobServiceServer
}

func Register() servicegrpc.MergeBlobServiceServer {
	return &grpcHandle{}
}
