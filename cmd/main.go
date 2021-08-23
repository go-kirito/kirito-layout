package main

import (
	"flag"
	"log"

	"app/di"

	"github.com/go-kirito/pkg/application"
	"github.com/go-kirito/pkg/middleware/recovery"
	"github.com/go-kirito/pkg/transport/grpc"
	"github.com/go-kirito/pkg/transport/http"
	"github.com/go-kirito/pkg/util/response"
	"github.com/go-kirito/pkg/zconfig"
	"github.com/go-kirito/pkg/zlog"
)

var config string

var Name = "demo"

func init() {
	flag.StringVar(&config, "f", "config.yaml", "config path, eg: -f config.yaml")
}

func main() {

	flag.Parse()

	if err := zconfig.Load(config); err != nil {
		log.Fatal("读取配置文件失败:", config)
	}

	zlog.Init()

	grpcAddress := zconfig.GetString("server.grpc.port")

	if grpcAddress == "" {
		grpcAddress = ":9000"
	}

	grpcSrv := grpc.NewServer(
		grpc.Address(grpcAddress),
		grpc.Middleware(
			recovery.Recovery(),
		),
	)

	httpAddress := zconfig.GetString("server.http.port")

	if httpAddress == "" {
		httpAddress = ":8000"
	}

	httpSrv := http.NewServer(
		http.Address(httpAddress),
		http.ResponseEncoder(response.Encoder),
		http.ErrorEncoder(response.ErrorEncoder),
		http.Middleware(
			recovery.Recovery(),
		),
	)

	app := application.New(
		application.Name(Name),
		application.GrpcServer(grpcSrv),
		application.HttpServer(httpSrv),
	)

	if err := di.RegisterService(app); err != nil {
		log.Fatal(err)
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
