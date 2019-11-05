/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-11 14:06:18
 * @LastEditTime: 2019-08-11 14:39:27
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/api"
	rapi "github.com/micro/go-micro/api/handler/api"
	"github.com/micro/go-micro/errors"
	proto "github.com/yuwe1/micolearn/day02/mico-api/meta/proto"
)

type Example struct{}
type Foo struct{}

// Call 方法在下面main中我们通过endpoint将其注册到/example/call路由
func (e *Example) Call(ctx context.Context, req *proto.CallRequest, rsp *proto.CallResponse) error {
	log.Print("Meta Example.Call接口收到请求")

	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.api.example", "no content")
	}

	rsp.Message = "Meta已经收到你的请求，" + req.Name
	return nil
}

// Bar 方法在下面main中我们通过endpoint将其注册到/example/foo/bar路由
func (f *Foo) Bar(ctx context.Context, req *proto.EmptyRequest, rsp *proto.EmptyResponse) error {
	log.Print("Meta Foo.Bar接口收到请求")
	// noop

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)
	// 注册Example接口处理器
	proto.RegisterExampleHandler(service.Server(), new(Example), api.WithEndpoint(&api.Endpoint{
		// 接口方法，一定要在proto接口中存在，不能是类的自有方法
		Name: "Example.Call",
		// http请求路由，支持POSIX风格
		Path: []string{"/example/call"},
		// 支持的方法类型
		Method: []string{"POST"},
		// 该接口使用的API转发模式
		Handler: rapi.Handler,
	}))

	// 注册Foo接口处理器
	proto.RegisterFooHandler(service.Server(), new(Foo), api.WithEndpoint(&api.Endpoint{
		Name:    "Foo.Bar",
		Path:    []string{"/foo/bar"},
		Method:  []string{"POST"},
		Handler: rapi.Handler,
	}))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
