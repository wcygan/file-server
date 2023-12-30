package client

import (
	"context"
	"fmt"
	pb "github.com/wcygan/fs/api/golang/file"
	"google.golang.org/grpc"
)

func Upload(filename string) error {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()

	ctx := context.TODO()

	client := pb.NewFileServiceClient(conn)
	stream, err := client.Upload(ctx)

	req := pb.FileUploadRequest{
		Filename: filename,
		Content:  []byte("hello"),
	}

	err = stream.Send(&req)
	if err != nil {
		return fmt.Errorf("failed to send: %v", err)
	}

	stream.CloseAndRecv()
	fmt.Println("Upload successful")
	return nil
}
