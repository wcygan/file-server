package server

import (
	"context"
	"fmt"
	pb "github.com/wcygan/fs/api/golang/file"
	//"io"
)

type Server struct {
	pb.UnimplementedFileServiceServer
	data map[string][]byte
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Upload(stream pb.FileService_UploadServer) error {
	//for {
	//	// Receive a message from the stream
	//	req, err := stream.Recv()
	//	if err == io.EOF {
	//		// End of the stream
	//		return stream.SendAndClose(&pb.FileUploadResponse{ /* response fields */ })
	//	}
	//	if err != nil {
	//		return err
	//	}
	//
	//	// Process the request (e.g., save file chunk)
	//	// ...
	//
	//}
	return nil
}

func (s *Server) Download(req *pb.FileDownloadRequest, stream pb.FileService_DownloadServer) error {
	//// Retrieve the file or data you want to send
	//// ...
	//
	//// Stream the file or data back to the client
	//for /* condition */ {
	//	// Prepare a chunk or part of the file/data
	//	response := &pb.FileDownloadResponse{ /* response fields */ }
	//
	//	// Send it to the client
	//	if err := stream.Send(response); err != nil {
	//		return err
	//	}
	//}
	//
	return nil
}

func (s *Server) Delete(ctx context.Context, req *pb.FileDeleteRequest) (*pb.FileDeleteResponse, error) {
	if s.data[req.GetFilename()] == nil {
		return &pb.FileDeleteResponse{Message: fmt.Sprintf("key '%s' does not exist", req.GetFilename())}, nil
	} else {
		delete(s.data, req.GetFilename())
		return &pb.FileDeleteResponse{Message: fmt.Sprintf("key '%s' deleted", req.GetFilename())}, nil
	}
}
