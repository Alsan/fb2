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

func GetBufferSize(fileSize int64) int {
	if fileSize < 1024 {
		return 1024
	}

	if fileSize < 1024*8 {
		return 8192
	}

	if fileSize < 1024*64 {
		return 1024 * 64
	}

	if fileSize < 1024*1024 {
		return 1024 * 1024
	}

	return 1024 * 1024 * 2
}
