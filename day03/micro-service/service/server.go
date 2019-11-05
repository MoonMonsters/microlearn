/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 10:58:08
 * @LastEditTime: 2019-08-15 11:17:50
 * @LastEditors: Please set LastEditors
 */
package main
import(
	"context"
	"fmt"
	proto "github.com/yuwe1/micolearn/day03/micro-service/proto"
	"github.com/micro/go-micro"
)
type Greeter struct{}
func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "你好，" + req.Name
	return nil
}
func main() {
	service :=micro.NewService(
		micro.Name("greeter.service"),
		micro.Version("last"),
		micro.Metadata(map[string]string{
			"type":"你好，世界",
		}),
	)
	service.Init()
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))
	// 启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
