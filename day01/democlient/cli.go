/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-09 09:53:36
 * @LastEditTime: 2019-08-10 17:41:47
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"context"
	"fmt"

	microclient "github.com/micro/go-micro/client"
	demo "microlearn/day01/demoservice/proto/demo"
)

const (
	address = "localhost:50051"
)

func main() {
	client := demo.NewSayService("go.micro.srv.demo", microclient.DefaultClient)
	r, err := client.Hello(context.Background(), &demo.Request{Msg: "hahaahah"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
