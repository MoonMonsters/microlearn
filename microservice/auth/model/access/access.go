/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-18 08:42:36
 * @LastEditTime: 2019-08-18 08:51:04
 * @LastEditors: Please set LastEditors
 */
package access

import(
	"fmt"
	"sync"
	r "github.com/go-redis/redis"
	"github.com/yuwe1/micolearn/microservice/basic/redis"
)
var (
	s *service
	ca *r.Client
	m  sync.RWMutex
)
type service struct{

}
// 用户服务类
type Service interface{
	// MakeAccessToken 生成token
	MakeAccessToken(subject *Subject) (ret string, err error)

	// GetCachedAccessToken 获取缓存的token
	GetCachedAccessToken(subject *Subject) (ret string, err error)

	// DelUserAccessToken 清除用户token
	DelUserAccessToken(token string) (err error)
}
// GetService 获取服务类
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

	ca = redis.GetRedis()

	s = &service{}
}