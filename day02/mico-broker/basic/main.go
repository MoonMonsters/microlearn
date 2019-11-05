/*
 * @Description: In User Settings Edit
 * @Author: 陌无崖
 * @Date: 2019-08-12 13:12:59
 * @LastEditTime: 2019-08-12 14:55:00
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"
	"time"

	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/util/log"
)

var (
	// 定义一个待发布的主题
	topic = "mu.micro.book.topic.payment.mowuya"
	b     broker.Broker
)

func pub() {
	// 该Ticker包含一个通道字段，并会每隔时间段d就向该通道发送当时的时间。
	// 它会调整时间间隔或者丢弃tick信息以适应反应慢的接收者。如果d<=0会panic。关闭该Ticker可以释放相关资源。
	tick := time.NewTicker(time.Second)
	i := 0
	for range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d:%s", i, time.Now().String())),
		}
		log.Infof(broker.String())
		// 发布消息
		if err := broker.Publish(topic, msg); err != nil {
			log.Infof("[pub] Message publication failed: %v", err)
		} else {
			fmt.Println("[pub] Message published: ", string(msg.Body))
		}
		i++
	}
}

func sub() {
	// 订阅消息并接收
	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		log.Info("[sub] Received Body: %s,Header: %s\n", string(p.Message().Body), p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	cmd.Init()
	if err := broker.Init(); err != nil {
		log.Fatalf("broker.Init() error: %v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("broker.Connect() error: %v", err)
	}
	go pub()
	go sub()
	// 20秒后过期，并向当前通道发送事件，结束主线程
	<-time.After(time.Second * 20)
}
