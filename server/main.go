package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	fb "github.com/filebrowser/filebrowser/proto"
	"github.com/filebrowser/filebrowser/utils"
	"github.com/filebrowser/filebrowser/v2/server/rpc"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

const PORT = 8080

func grpcServer(listener net.Listener) {
	s := grpc.NewServer()
	fb.RegisterFileBrowserRpcServiceServer(s, &rpc.Server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Unable to start grpc server: %v", err)
	}
}

func httpServer(listener net.Listener) {
	// the handler HAVE TO BE registered before Serve() is called
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	if err := http.Serve(listener, nil); err != nil {
		log.Fatalf("Unable to start http server: %v", err)
	}
}

func main() {
	// create a listener at the desired port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		utils.ExitIfError(fmt.Sprintf("Unable to listen to port :%d", PORT), err)
	}

	// close the listener when done
	defer listener.Close()

	// create a cmux object
	tcpm := cmux.New(listener)

	// declare the match rules for different services requested
	grpcFilter := tcpm.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpFilter := tcpm.Match(cmux.HTTP1Fast())

	// initialize the servers by passing in the custom listeners
	go grpcServer(grpcFilter)
	go httpServer(httpFilter)

	log.Printf("server listenering on port :%d", PORT)
	if err := tcpm.Serve(); err != nil {
		log.Fatalf("Error serving cmux: %v", err)
	}
}
