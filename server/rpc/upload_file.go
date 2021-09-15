package rpc

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"

	c "github.com/alsan/fb2/common"
	fb "github.com/alsan/fb2/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UploadFile(stream fb.FileBrowserRpcService_UploadFileServer) error {
	req, err := stream.Recv()
	if err != nil {
		return logError(err)
	}

	metaData := req.GetMetadata()
	if !verifyToken(stream.Context(), metaData.GetToken()) {
		return logError(status.Errorf(codes.PermissionDenied, "unable to verify token"))
	}

	filename := filepath.Join(serverConf.Root, metaData.GetPath(), metaData.GetFilename())
	filepath := filepath.Dir(filename)

	// ensure file path is exist
	if err := os.MkdirAll(filepath, os.ModePerm); err != nil {
		return logError(status.Errorf(codes.Internal, "error creating directory: %v", err))
	}

	// create file handler
	file, err := os.Create(filename)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "error creating file: %v", err))
	}
	defer file.Close()

	// create buffered writter
	writter := bufio.NewWriter(file)

	for {
		// check grpc context
		if err := contextError(stream.Context()); err != nil {
			return err
		}

		// begin receive stream
		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive stream request: %v", err))
		}

		// write received chunk into file buffer
		if _, err := writter.Write(req.GetContent()); err != nil {
			return logError(status.Errorf(codes.Internal, "error writting chunk: %v", err))
		}
	}

	// flush file buffer
	writter.Flush()

	// validating written file size
	written := c.GetFileSize(file)
	if written != metaData.GetSize() {
		return logError(status.Errorf(codes.Unknown, "written size mismatch: %d", writter.Available()))
	}

	// validating file checksum
	checksum := c.GetFileChecksum(file)
	declared := metaData.GetChecksum()
	if checksum != declared {
		return logError(status.Errorf(codes.Unknown, "file checksum mismatch: %s, declared: %s", checksum, declared))
	}

	// reply client everything is ok
	if err = stream.SendAndClose(&fb.UploadFileReply{Status: fb.ReplyStatus_Ok}); err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot send response, %v", err))
	}

	return nil
}
