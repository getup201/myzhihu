package main

import (
	"flag"
	"fmt"

	"myzhihu/apps/app/internal/config"
	"myzhihu/apps/app/internal/handler"
	"myzhihu/apps/app/internal/svc"
	"myzhihu/pkg/xcode"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/applet-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c) // 加载配置文件并解析到结构体 yaml文件是传到了这里

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 注册自定义错误处理方法   唯一增加的地方
	//gozero 提供的自定义错误处理方法  把自定义的方法传进来
	httpx.SetErrorHandler(xcode.ErrHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

//发文章的话只要前两个服务
//可以先只实现前两个微服务
// user
// article
// like
// follow
// Message

//下面三个功能可以砍掉
// member
// qa
// chat
