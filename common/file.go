package common

import "os"

func GetFileSize(file *os.File) int64 {
	fi, err := file.Stat()
	CheckErr(err)

	return fi.Size()
}
