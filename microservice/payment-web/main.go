/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-20 09:49:14
 * @LastEditTime: 2019-08-24 14:35:59
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"
	"net/http"

	"github.com/yuwe1/micolearn/microservice/basic"
	"github.com/yuwe1/micolearn/microservice/basic/common"
	config "github.com/yuwe1/micolearn/microservice/basic/gconfig"
	"github.com/micro/go-plugins/config/source/grpc"
	"github.com/yuwe1/micolearn/microservice/payment-web/handler"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
)
var (
	appName = "payment_web"
	cfg     = &appCfg{}
)

type appCfg struct {
	common.AppCfg
}
func main() {
	// 初始化配置
	initCfg()

	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	// 创建新服务
	service := web.NewService(
		web.Name(cfg.Name),
		web.Version(cfg.Version),
		web.Registry(micReg),
		web.Address(cfg.Addr()),
	)

	// 初始化服务
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	// 新建订单接口
	authHandler := http.HandlerFunc(handler.PayOrder)
	service.Handle("/payment/pay-order", handler.AuthWrapper(authHandler))

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := &common.Consul{}
	err := config.C().App("consul", consulCfg)
	if err != nil {
		panic(err)
	}

	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.Host, consulCfg.Port)}
}

func initCfg() {
	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)

	basic.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	log.Logf("[initCfg] 配置，cfg：%v", cfg)

	return
}