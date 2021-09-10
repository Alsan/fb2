module github.com/alsan/filebrowser/client

go 1.17

require (
	github.com/alsan/filebrowser/common v0.0.0
	github.com/alsan/filebrowser/proto v0.0.0
)

require (
	github.com/spf13/cobra v1.2.1
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	google.golang.org/grpc v1.40.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/alsan/filebrowser/proto v0.0.0 => ../proto

replace github.com/alsan/filebrowser/common v0.0.0 => ../common
