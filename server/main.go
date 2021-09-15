package main

import (
	"runtime"

	"github.com/alsan/fb2/server/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
