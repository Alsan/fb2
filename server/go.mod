module github.com/filebrowser/filebrowser/v2/server

go 1.17

require (
	github.com/filebrowser/filebrowser/proto v0.0.0
	github.com/filebrowser/filebrowser/utils v0.0.0
	github.com/soheilhy/cmux v0.1.5
	google.golang.org/grpc v1.40.0
)

require (
	github.com/golang/protobuf v1.5.0 // indirect
	golang.org/x/net v0.0.0-20201202161906-c7110b5ffcbb // indirect
	golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/filebrowser/filebrowser/proto v0.0.0 => ../proto

replace github.com/filebrowser/filebrowser/utils v0.0.0 => ../utils
