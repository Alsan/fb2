package rpc

import (
	"log"

	fb "github.com/alsan/filebrowser/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UploadFile(stream fb.FileBrowserRpcService_UploadFileServer) error {
	log.Println("upload file request")

	req, err := stream.Recv()
	if err != nil {
		return logError(err)
	}

	token := req.GetToken()
	if verifyToken(stream.Context(), token) {
		return nil
	}

	return logError(status.Errorf(codes.Unknown, tokenTimeoutMessage))
}
