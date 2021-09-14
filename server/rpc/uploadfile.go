package rpc

import (
	"bytes"
	"io"
	"log"

	fb "github.com/alsan/filebrowser/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UploadFile(stream fb.FileBrowserRpcService_UploadFileServer) error {
	req, err := stream.Recv()
	if err != nil {
		return logError(err)
	}

	metaData := req.GetMetadata()
	if verifyToken(stream.Context(), metaData.GetToken()) {
		buffer := bytes.Buffer{}
		bufSize := 0

		for {
			if err := contextError(stream.Context()); err != nil {
				return err
			}

			req, err := stream.Recv()
			if err == io.EOF {
				log.Print("no more data")
				break
			}
			if err != nil {
				return logError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
			}

			chunk := req.GetContent()
			size := len(chunk)
			bufSize += size

			if _, err = buffer.Write(chunk); err != nil {
				return logError(status.Errorf(codes.Unknown, "cannot write chunk data, %v", err))
			}
		}

		if err = stream.SendAndClose(&fb.UploadFileReply{Status: fb.ReplyStatus_Ok}); err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot send response, %v", err))
		}

		return nil
	}

	return logError(status.Errorf(codes.Unknown, tokenTimeoutMessage))
}
