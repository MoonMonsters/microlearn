/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-10 15:22:34
 * @LastEditTime: 2019-08-11 10:55:18
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	api "github.com/yuwe1/micolearn/day02/mico-api/api/proto"
)

type Example struct{}

type Foo struct{}

// Example.Call 通过API向外暴露为/example/call，接收http请求
// 即：/example/call请求会调用go.micro.api.example服务的Example.Call方法
func (e *Example) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Example.Call接口收到请求")

	name, ok := req.Get["name"]

	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.example", "参数不正确")
	}

	// 打印请求头
	for k, v := range req.Header {
		log.Log("请求头信息，", k, " : ", v)
	}

	rsp.StatusCode = 200

	b, _ := json.Marshal(map[string]string{
		"message": "我们已经收到你的请求，" + strings.Join(name.Values, " "),
	})

	// 设置返回值
	rsp.Body = string(b)

	return nil
}

// Bar 方法全称是Foo.Bar，故而它会以/example/foo/bar为路由提供服务
//
func (f *Foo) Bar(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Logf("Foo.Bar接口收到请求")

	if req.Method != "POST" {
		return errors.BadRequest("go.micro.api.example", "require post")
	}

	ct, ok := req.Header["Content-Type"]
	if !ok || len(ct.Values) == 0 {
		return errors.BadRequest("go.micro.api.example", "need content-type")
	}

	fmt.Println(ct.Values[0])
	if ct.Values[0] != "application/json" {
		return errors.BadRequest("go.micro.api.example", "expect application/json")
	}

	var body map[string]interface{}
	json.Unmarshal([]byte(req.Body), &body)

	// 设置返回值
	rsp.Body = "收到消息：" + string([]byte(req.Body))

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"),
		micro.Version("latest"),
	)

	service.Init()
	// 注册 example handler
	api.RegisterExampleHandler(service.Server(), new(Example))

	// 注册 foo handler
	api.RegisterFooHandler(service.Server(), new(Foo))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
