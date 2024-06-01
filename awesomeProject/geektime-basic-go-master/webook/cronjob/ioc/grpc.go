package ioc

import (
	grpc2 "gitee.com/geekbang/basic-go/webook/cronjob/grpc"
	"gitee.com/geekbang/basic-go/webook/pkg/grpcx"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func InitGRPCxServer(cronJobGrpc *grpc2.CronJobServiceServer) *grpcx.Server {
	type Config struct {
		Addr string `yaml:"addr"`
	}
	var cfg Config
	err := viper.UnmarshalKey("grpc.server", &cfg)
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	cronJobGrpc.Register(server)
	return &grpcx.Server{
		Server: server,
		Addr:   cfg.Addr,
	}
}
