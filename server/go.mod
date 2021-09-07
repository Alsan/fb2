module github.com/alsan/filebrowser/server

go 1.17

require (
	github.com/DataDog/zstd v1.4.8 // indirect
	github.com/Sereal/Sereal v0.0.0-20210713121911-8c71d8dbe594 // indirect
	github.com/alsan/filebrowser/common v0.0.0
	github.com/alsan/filebrowser/proto v0.0.0
	github.com/asdine/storm v2.1.2+incompatible
	github.com/caddyserver/caddy v1.0.3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/disintegration/imaging v1.6.2
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/dsoprea/go-exif/v3 v3.0.0-20201216222538-db167117f483
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/websocket v1.4.1
	github.com/maruel/natural v0.0.0-20180416170133-dbcb3e2e8cf1
	github.com/marusama/semaphore/v2 v2.4.1
	github.com/mholt/archiver v3.1.1+incompatible
	github.com/mitchellh/go-homedir v1.1.0
	github.com/nwaples/rardecode v1.1.2 // indirect
	github.com/pelletier/go-toml v1.6.0
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/rs/zerolog v1.24.0
	github.com/soheilhy/cmux v0.1.4
	github.com/spf13/afero v1.2.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.1
	github.com/stretchr/testify v1.6.1
	github.com/tomasen/realip v0.0.0-20180522021738-f0c99a92ddce
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	go.etcd.io/bbolt v1.3.3
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/appengine v1.5.0 // indirect
	google.golang.org/grpc v1.40.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.3.0
)

require (
	github.com/cenkalti/backoff v2.1.1+incompatible // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dsoprea/go-logging v0.0.0-20200710184922-b02d349568dd // indirect
	github.com/dsoprea/go-utility/v2 v2.0.0-20200717064901-2fccff4aa15e // indirect
	github.com/flynn/go-shlex v0.0.0-20150515145356-3f9db97f8568 // indirect
	github.com/fsnotify/fsnotify v1.4.7 // indirect
	github.com/go-acme/lego v2.5.0+incompatible // indirect
	github.com/go-errors/errors v1.1.1 // indirect
	github.com/golang/geo v0.0.0-20200319012246-673a6f80352d // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/klauspost/cpuid v1.2.0 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mholt/certmagic v0.6.2-0.20190624175158-6a42ef9fe8c2 // indirect
	github.com/miekg/dns v1.1.3 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/cast v1.3.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/ulikunitz/xz v0.5.10 // indirect
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98 // indirect
	google.golang.org/grpc/examples v0.0.0-20210903175933-b2ba77a36ff8 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/ini.v1 v1.51.0 // indirect
	gopkg.in/square/go-jose.v2 v2.2.2 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/alsan/filebrowser/proto v0.0.0 => ../proto

replace github.com/alsan/filebrowser/common v0.0.0 => ../common
