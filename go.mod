module airplaneServer

// go version
go 1.14

// gin framework
require github.com/gin-gonic/gin v1.6.3

// log plugin
require github.com/sirupsen/logrus v1.6.0

// gorm
require github.com/jinzhu/gorm v1.9.14

// env plugin
require github.com/joho/godotenv v1.3.0

// redis
require github.com/go-redis/redis/v8 v8.2.3

require (
	github.com/DataDog/sketches-go v0.0.0-20190923095040-43f19ad77ff7 // indirect
	github.com/benbjohnson/clock v1.0.3 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/opentracing/opentracing-go v1.1.1-0.20190913142402-a7454ce5950e // indirect
	// testing plugin
	github.com/stretchr/testify v1.6.1
	// DI plugin
	go.uber.org/dig v1.10.0
	golang.org/x/exp v0.0.0-20200513190911-00229845015e // indirect
	google.golang.org/genproto v0.0.0-20191009194640-548a555dbc03 // indirect
	google.golang.org/grpc v1.30.0 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)
