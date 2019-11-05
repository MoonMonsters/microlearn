/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-11 10:56:47
 * @LastEditTime: 2019-08-11 11:12:13
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"context"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	proto "github.com/yuwe1/micolearn/day02/mico-api/api/proto"
)

// 切记，事件订阅结构的所有公有方法都会被执行 方法名没有限制，但是方法一定要接收ctx，event
type Event struct {
}

func (e *Event) Process(ctx context.Context, event *proto.Event) error {
	log.Log("公有方法Process 收到事件，", event.Name)
	log.Log("公有方法Process 数据", event.Data)
	return nil
}
func (e *Event) Process2(ctx context.Context, event *proto.Event) error {
	log.Log("公有方法Process2 收到事件，", event.Name)
	log.Log("公有方法Process2 数据", event.Data)
	return nil
}
func (e *Event) process(ctx context.Context, event *proto.Event) error {
	log.Log("私有方法process，收到事件，", event.Name)
	log.Log("私有方法process，数据", event.Data)
	return nil
}
func main() {
	service := micro.NewService(
		micro.Name("go.micro.evt.user"),
	)
	service.Init()
	micro.RegisterSubscriber("go.micro.evt.user", service.Server(), new(Event))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
