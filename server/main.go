package main

import (
	"runtime"

	"github.com/filebrowser/filebrowser/server/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
