module github.com/filebrowser/filebrowser/server

go 1.17

require (
	github.com/filebrowser/filebrowser/proto v0.0.0
	github.com/filebrowser/filebrowser/utils v0.0.0
	github.com/soheilhy/cmux v0.1.5
	google.golang.org/grpc v1.40.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/spf13/cobra v1.2.1
	golang.org/x/mod v0.5.0
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

require (
	github.com/asdine/storm v1.1.0
	github.com/caddyserver/caddy v1.0.5
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/disintegration/imaging v1.6.2
	github.com/dsoprea/go-exif/v3 v3.0.0-20210625224831-a6301f85c82b
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.4.2
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/maruel/natural v1.0.0
	github.com/marusama/semaphore/v2 v2.5.0
	github.com/mholt/archiver v2.1.0+incompatible
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/pelletier/go-toml v1.9.3
	github.com/rs/zerolog v1.24.0
	github.com/spf13/afero v1.6.0
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.8.1
	github.com/stretchr/testify v1.7.0
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/tomasen/realip v0.0.0-20180522021738-f0c99a92ddce
	go.etcd.io/bbolt v1.3.6
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/DataDog/zstd v1.4.8 // indirect
	github.com/Sereal/Sereal v0.0.0-20210713121911-8c71d8dbe594 // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/dsoprea/go-logging v0.0.0-20200710184922-b02d349568dd // indirect
	github.com/dsoprea/go-utility/v2 v2.0.0-20200717064901-2fccff4aa15e // indirect
	github.com/flynn/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/go-errors/errors v1.1.1 // indirect
	github.com/golang/geo v0.0.0-20200319012246-673a6f80352d // indirect
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/nwaples/rardecode v1.1.2 // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/ulikunitz/xz v0.5.10 // indirect
	gopkg.in/vmihailenco/msgpack.v2 v2.9.2 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/filebrowser/filebrowser/proto v0.0.0 => ../proto

replace github.com/filebrowser/filebrowser/utils v0.0.0 => ../utils
