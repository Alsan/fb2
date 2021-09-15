package rpc

import (
	"bufio"
	"io"
	"os"
	"path/filepath"

	c "github.com/alsan/filebrowser/common"
	fb "github.com/alsan/filebrowser/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DownloadFile(in *fb.DownloadFileRequest, stream fb.FileBrowserRpcService_DownloadFileServer) error {
	if !verifyToken(stream.Context(), in.GetToken()) {
		return logError(status.Errorf(codes.PermissionDenied, "unable to verify token"))
	}

	filename := filepath.Join(serverConf.Root, in.GetFilename())
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return logError(status.Errorf(codes.PermissionDenied, "unable to open file"))
	}

	fileSize := c.GetFileSize(file)

	if err := stream.Send(&fb.DownloadFileReply{
		Status: fb.ReplyStatus_Ok,
		Reply: &fb.DownloadFileReply_Data{
			Data: &fb.DownloadFileData{
				Option: &fb.DownloadFileData_Info{
					Info: &fb.DownloadFileInfo{
						Filename: filepath.Base(filename),
						Size:     fileSize,
						Checksum: c.GetFileChecksum(file),
					},
				},
			},
		},
	}); err != nil {
		return logError(status.Errorf(codes.Unknown, "unable to send file info"))
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, c.GetBufferSize(fileSize))

	file.Seek(0, io.SeekStart)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Internal, "Unable to reader chunk, %v", err))
		}

		if err = stream.Send(&fb.DownloadFileReply{
			Status: fb.ReplyStatus_Ok,
			Reply: &fb.DownloadFileReply_Data{
				Data: &fb.DownloadFileData{
					Option: &fb.DownloadFileData_Content{
						Content: buffer[:n],
					},
				},
			},
		}); err != nil {
			return logError(status.Errorf(codes.Unknown, "unable to send chunk, %s", err))
		}
	}

	return nil
}
