/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-16 09:05:11
 * @LastEditTime: 2019-08-16 10:11:06
 * @LastEditors: Please set LastEditors
 */
// 创建一个发布消息的服务
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
)

var (
	topic = "go.micro.web.topic.hi"
)

func pub(name string) {
	msg := &broker.Message{
		Header: map[string]string{
			"name": fmt.Sprintf("%s", name),
		},
		Body: []byte(fmt.Sprintf("%s:%s", name, time.Now().String())),
	}
	if err := broker.Publish(topic, msg); err != nil {
		log.Logf("[pub] 发布消息：%s", err)
	} else {
		log.Logf("[pub] 发布消息: %s", string(msg.Body))
	}
}

// 定义一个handler
func hi(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	_ = r.ParseForm()
	// 返回结果
	response := map[string]interface{}{
		"ref":  time.Now().UnixNano(),
		"data": "Hello!" + r.Form.Get("name"),
	}
	//返回json结构
	// NewEncoder创建一个将数据写入w的*Encoder。
	// Encode将v的json编码写入输出流，并会写入一个换行符，
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	pub(r.Form.Get("name"))
}
func main() {
	// 创建一个服务
	service := web.NewService(
		web.Name("go.micro.book.web.pub"),
		web.Version("latest"),
		web.Address(":8088"),
	)
	// 初始化我们的服务
	_ = service.Init()
	service.HandleFunc("/hi", hi)
	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
