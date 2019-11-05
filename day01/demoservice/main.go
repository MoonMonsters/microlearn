/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-09 09:32:45
 * @LastEditTime: 2019-08-10 17:34:35
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"

	demo "microlearn/day01/demoservice/proto/demo"
)

const (
	port = ":50051"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *demo.Request, rsp *demo.Response) error {
	fmt.Println("received", req.Msg)
	rsp.Header = make(map[string]*demo.Pair)
	rsp.Header["name"] = &demo.Pair{Key: 1, Values: "abc"}
	rsp.Msg = "Hello World"
	rsp.Values = append(rsp.Values, "a", "b")
	rsp.Type = demo.RespType_DESCEND
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.demo"),
		micro.Version("latest"),
	)
	srv.Init()
	demo.RegisterSayHandler(srv.Server(), new(Say))
	// run server
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
