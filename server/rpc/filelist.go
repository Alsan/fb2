package rpc

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"time"

	c "github.com/alsan/filebrowser/common"
	fb "github.com/alsan/filebrowser/proto"
)

func getFilesWithoutFilter(path string) []string {
	var files []string

	err := filepath.Walk(serverConf.Root+path, func(curr string, info os.FileInfo, err error) error {
		files = append(files, curr)

		return nil
	})
	if err != nil {
		panic(err)
	}

	return files
}

func getFilesWithFilter(path, filter string) []string {
	var files []string
	var filterList = c.StrSlice(strings.Split(filter, ","))

	err := filepath.Walk(serverConf.Root+path, func(curr string, info os.FileInfo, err error) error {
		if filterList.Has(filepath.Ext(curr)) {
			files = append(files, curr)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	return files

}

func getFileList(path, filter string) []string {
	if filter == "" {
		return getFilesWithoutFilter(path)
	}

	return getFilesWithFilter(path, filter)
}

func (s *Server) FileList(ctx context.Context, in *fb.FileListRequest) (*fb.FileListReply, error) {
	token := in.GetToken()
	ip := getRemoteIP(ctx)
	userToken := token + "-" + ip
	timestamp, ok := session[userToken]
	now := time.Now().Unix()

	if ok && now-timestamp < int64(time.Minute)*5 {
		// extend user sesion
		session[userToken] = now

		// return the file list
		return &fb.FileListReply{
			Status: fb.ReplyStatus_Ok,
			Data: &fb.FileListReply_List{
				List: &fb.List{
					Item: getFileList(in.GetPath(), in.GetFilter()),
				},
			},
		}, nil
	}

	// timeout or get the token from another ip
	return &fb.FileListReply{
		Status: fb.ReplyStatus_Failed,
		Data: &fb.FileListReply_Message{
			Message: "Invalid token, please login first",
		},
	}, nil
}
