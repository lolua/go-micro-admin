module micro-admin

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis/v8 v8.0.0-beta.2
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.7.0
	github.com/prometheus/common v0.6.0
	google.golang.org/protobuf v1.24.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
