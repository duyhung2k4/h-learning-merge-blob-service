package grpchandle

import "app/generated/grpc/servicegrpc"

type grpcHandle struct {
	servicegrpc.UnimplementedMergeBlobServiceServer
}

func Register() servicegrpc.MergeBlobServiceServer {
	return &grpcHandle{}
}
