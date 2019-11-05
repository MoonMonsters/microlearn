/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-17 08:48:03
 * @LastEditTime: 2019-08-17 09:50:45
 * @LastEditors: Please set LastEditors
 */
package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"
	us "github.com/yuwe1/micolearn/microservice/user-srv/model/user"
	s "github.com/yuwe1/micolearn/microservice/user-srv/proto/user"
)

type Service struct{}

var (
	userService us.Service
)

// 初始化handler
func Init() {
	var err error
	userService, err = us.GetService()
	if err != nil {
		log.Fatal("handler初始化错误")
	}

}

// QueryUserByName 通过参数中的名字返回用户
func (e *Service) QueryUserByName(ctx context.Context, req *s.Request, rsp *s.Response) error {
	user, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Success = false
		rsp.Error = &s.Error{
			Code:   500,
			Detail: err.Error(),
		}

		return err
	}

	rsp.User = user
	rsp.Success = true
	return nil
}
