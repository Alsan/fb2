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
	log.Println("upload file request")

	req, err := stream.Recv()
	if err != nil {
		return logError(err)
	}
	log.Println("steam receiving")

	log.Println("getting metadata")
	metaData := req.GetMetadata()
	if verifyToken(stream.Context(), metaData.GetToken()) {
		log.Println("token verified")
		buffer := bytes.Buffer{}
		bufSize := 0

		for {
			log.Println("checking context error")
			err := contextError(stream.Context())
			if err != nil {
				return err
			}

			log.Print("waiting to receive more data")

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
			log.Printf("received a chunk with size: %d bytes", size)
			bufSize += size

			_, err = buffer.Write(chunk)
			if err != nil {
				return logError(status.Errorf(codes.Unknown, "cannot write chunk data, %v", err))
			}

			res := &fb.UploadFileReply{
				Status: fb.ReplyStatus_Ok,
			}

			err = stream.SendAndClose(res)
			if err != nil {
				return logError(status.Errorf(codes.Unknown, "cannot send response, %v", err))
			}
		}
		return nil
	}

	return logError(status.Errorf(codes.Unknown, tokenTimeoutMessage))
}
