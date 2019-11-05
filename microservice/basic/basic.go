/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-16 17:37:00
 * @LastEditTime: 2019-08-24 14:04:44
 * @LastEditors: Please set LastEditors
 */
package basic

// import (
// 	"github.com/yuwe1/micolearn/microservice/basic/config"
// 	"github.com/yuwe1/micolearn/microservice/basic/db"
// 	"github.com/yuwe1/micolearn/microservice/basic/redis"
// )
// func Init(){
// 	config.Init()
// 	db.Init()
// 	redis.Init()
// }
import (
	config "github.com/yuwe1/micolearn/microservice/basic/gconfig"
)

var (
	pluginFuncs []func()
)

type Options struct {
	EnableDB    bool
	EnableRedis bool
	cfgOps      []config.Option
}

type Option func(o *Options)

func Init(opts ...config.Option) {
	// 初始化配置
	config.Init(opts...)

	// 加载依赖配置的插件
	for _, f := range pluginFuncs {
		f()
	}
}

func Register(f func()) {
	pluginFuncs = append(pluginFuncs, f)
}