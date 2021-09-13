package common

import "os"

func GetFileSize(file *os.File) int64 {
	fi, err := file.Stat()
	CheckErr(err)

	return fi.Size()
}

func IsFileExist(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}

	return false
}
