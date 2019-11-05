/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-17 08:49:07
 * @LastEditTime: 2019-08-18 12:33:09
 * @LastEditors: Please set LastEditors
 */
package user

import (
	"fmt"
	"sync"

	proto "github.com/yuwe1/micolearn/microservice/user-srv/proto/user"
)

var (
	s *service
	m sync.RWMutex
)

// service服务
type service struct {
}

// 用户服务接口
type Service interface {
	QueryUserByName(userName string) (ret *proto.User, err error)
}

// 获取用户服务的接口
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init 初始化用户服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	s = &service{}
}
