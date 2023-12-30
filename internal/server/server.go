package server

import (
	"bytes"
	"context"
	"fmt"
	pb "github.com/wcygan/fs/api/golang/file"
	"io"
	"log"
	"sync"
)

type Server struct {
	pb.UnimplementedFileServiceServer
	files   map[string][]byte
	filesMu sync.RWMutex
}

func NewServer() *Server {
	return &Server{}
}

// Upload handles streaming file uploads
func (s *Server) Upload(stream pb.FileService_UploadServer) error {
	log.Printf("got a file upload request")
	var buffer bytes.Buffer
	var filename string
	isFirstMessage := true

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// End of stream
			break
		}
		if err != nil {
			return err
		}

		if isFirstMessage {
			filename = req.Filename
			isFirstMessage = false
		}

		buffer.Write(req.Content)
	}

	// Store the file in the server's memory
	s.filesMu.Lock()
	s.files[filename] = buffer.Bytes()

	for k, v := range s.files {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}

	s.filesMu.Unlock()

	return stream.SendAndClose(&pb.FileUploadResponse{
		Message: "File uploaded successfully",
	})
}

// Download handles streaming file downloads
func (s *Server) Download(req *pb.FileDownloadRequest, stream pb.FileService_DownloadServer) error {
	s.filesMu.RLock()
	fileContent, exists := s.files[req.Filename]
	s.filesMu.RUnlock()

	if !exists {
		return stream.Send(&pb.FileDownloadResponse{
			Content: nil,
		})
	}

	// Stream the file content
	for i := 0; i < len(fileContent); i += 1024 {
		end := i + 1024
		if end > len(fileContent) {
			end = len(fileContent)
		}

		if err := stream.Send(&pb.FileDownloadResponse{
			Content: fileContent[i:end],
		}); err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) Delete(ctx context.Context, req *pb.FileDeleteRequest) (*pb.FileDeleteResponse, error) {
	log.Printf("got a delete request for '%s'", req.GetFilename())
	if s.files[req.GetFilename()] == nil {
		return &pb.FileDeleteResponse{Message: fmt.Sprintf("key '%s' does not exist", req.GetFilename())}, nil
	} else {
		delete(s.files, req.GetFilename())
		return &pb.FileDeleteResponse{Message: fmt.Sprintf("key '%s' deleted", req.GetFilename())}, nil
	}
}
