package grpchandle

import (
	"app/generated/proto/servicegrpc"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (h *grpcHandle) SendMergeBlob(stream grpc.ClientStreamingServer[servicegrpc.SendMergeBlobRequest, servicegrpc.SendMergeBlobResponse]) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("metadata nil")
	}
	uuid := md["uuid"][0]

	readPipe, writePipe, err := os.Pipe()
	if err != nil {
		log.Fatalf("Error creating pipe: %v", err)
	}

	// Run FFmpeg play HLS
	go runFFmpeg(readPipe, uuid)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return errors.New("EOF grpc merge-blob-service")
		}

		if err != nil {
			return err
		}

		// Ghi dữ liệu nhận được vào pipe
		_, err = writePipe.Write(req.Blob)
		if err != nil {
			log.Printf("Error writing to pipe: %v", err)
		}
	}
}

func runFFmpeg(input *os.File, uuid string) {
	// path save HLS output
	outputPath := fmt.Sprintf("cmd/merge-blob/data/stream/%s", uuid)
	os.MkdirAll(outputPath, os.ModePerm)

	cmd := exec.Command("ffmpeg",
		"-f", "mpegts",
		"-i", "pipe:0", // Nhận từ pipe
		"-c", "copy", // Copy codec hiện tại, không mã hóa lại
		"-hls_time", "5", // Mỗi đoạn dài 1 giây
		"-hls_list_size", "0", // Lưu tất cả danh sách (không ghi đè)
		"-f", "hls",
		outputPath+"/stream.m3u8", // Đường dẫn lưu HLS
	)

	cmd.Stdin = input
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error running FFmpeg: %v", err)
	}
}
