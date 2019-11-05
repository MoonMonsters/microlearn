/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-16 09:19:21
 * @LastEditTime: 2019-08-16 09:39:57
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
)

var (
	topic = "go.micro.web.topic.hi"
)

func main() {
	bk := broker.NewBroker(
		broker.Addrs(fmt.Sprintf("%s:%d", "192.168.10.150", 11089)),
	)
	_, err := bk.Subscribe(topic, func(p broker.Event) error {
		log.Logf("[sub]:Received Body: %s,Header:%s", string(p.Message().Body), p.Message().Header)
		return nil
	})
	if err != nil {
		log.Logf("[ERR] err:%s", err)
	}
	s := micro.NewService(
		micro.Name("go.micro.book.web.sub"),
		micro.Version("latest"),
		micro.Address(":8099"),
		micro.Broker(bk),
	)
	s.Init()
	_ = s.Run()
}
