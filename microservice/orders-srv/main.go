/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-19 15:22:38
 * @LastEditTime: 2019-08-24 14:37:42
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/yuwe1/micolearn/microservice/basic"
	"github.com/yuwe1/micolearn/microservice/basic/common"
	config "github.com/yuwe1/micolearn/microservice/basic/gconfig"
	"github.com/micro/go-plugins/config/source/grpc"
	"github.com/yuwe1/micolearn/microservice/orders-srv/handler"
	"github.com/yuwe1/micolearn/microservice/orders-srv/model"
	proto "github.com/yuwe1/micolearn/microservice/orders-srv/proto/orders"
	"github.com/yuwe1/micolearn/microservice/orders-srv/subscriber"
)
var (
	appName = "orders_srv"
	cfg     = &appCfg{}
)
type appCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置、数据库等信息
	initCfg()
	// consul注册
	micReg := consul.NewRegistry(registryOptions)
	// 新建服务
	// 新建服务
	service := micro.NewService(
		micro.Name(cfg.Name),
		micro.Registry(micReg),
		micro.Version(cfg.Version),
		micro.Address(cfg.Addr()),
	)
	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			model.Init()
			handler.Init()
			subscriber.Init()
		}),
	)
	// 侦听订单支付消息
	err := micro.RegisterSubscriber(common.TopicPaymentDone, service.Server(), subscriber.PayOrder)
	if err != nil {
		log.Fatal(err)
	}
	// 注册服务
	proto.RegisterOrdersHandler(service.Server(), new(handler.Orders))

	// 启动服务
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