package main

import (
	"flag"
	"fmt"

	"myzhihu/apps/user/rpc/internal/config"
	"myzhihu/apps/user/rpc/internal/server"
	"myzhihu/apps/user/rpc/internal/svc"
	"myzhihu/apps/user/rpc/service"
	"myzhihu/pkg/interceptors"

	"github.com/zeromicro/go-zero/core/conf"
	cs "github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

// "myzhihu/apps/user/rpc/service"  和 github.com/zeromicro/go-zero/core/service 重名了 要给其中一个加一个别名
func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		service.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == cs.DevMode || c.Mode == cs.TestMode {
			reflection.Register(grpcServer)
		}
	})

	// 将 自定义拦截器 interceptors.ServerErrorInterceptor() 注册进来
	//这个是服务端的拦截器  将自定义Xcode转换成 rpc能识别的码
	//因为这个要通过rpc传输给api
	s.AddUnaryInterceptors(interceptors.ServerErrorInterceptor())

	defer s.Stop()

	// // 服务注册 后面再实现
	// err := consul.Register(c.Consul, fmt.Sprintf("%s:%d", c.ServiceConf.Prometheus.Host, c.ServiceConf.Prometheus.Port))
	// if err != nil {
	// 	fmt.Printf("register consul error: %v\n", err)
	// }

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
