/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-11 16:22:32
 * @LastEditTime: 2019-08-12 09:15:12
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	proto "github.com/yuwe1/micolearn/day02/mico-api/rpc/proto"
)

type Example struct {
}
type Foo struct {
}

func (e *Example) Call(ctx context.Context, req *proto.CallRequest, rsp *proto.CallResponse) error {
	log.Print("收到了Example.Call请求")
	if len(req.Name) == 0 {
		return errors.BadRequest("go.micro.api.example", "no content")
	}
	rsp.Message = "RPC Call收到了你的请求 " + req.Name
	return nil
}
func (f *Foo) Bar(ctx context.Context, req *proto.EmptyRequest, rsp *proto.EmptyResponse) error {
	log.Print("收到 Foo.Bar 请求")
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
	)
	service.Init()
	// 注册example接口
	proto.RegisterExampleHandler(service.Server(), new(Example))
	// 注册foo接口
	proto.RegisterFooHandler(service.Server(), new(Foo))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
