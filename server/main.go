package main

import (
	"runtime"

	"github.com/alsan/filebrowser/server/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
