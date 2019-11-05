/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-16 18:31:28
 * @LastEditTime: 2019-08-18 12:31:14
 * @LastEditors: Please set LastEditors
 */
package db

import (
	"database/sql"
	"fmt"

	"github.com/micro/go-micro/util/log"
	"github.com/yuwe1/micolearn/microservice/basic/config"
)

func Init() {
	m.Lock()
	defer m.Unlock()
	var err error
	if inited {
		err = fmt.Errorf("[Init] db 已经初始化过")
		log.Logf(err.Error())
		return
	}
	if config.GetMysqlConfig().GetEnabled() {
		initMysql()
	}
	inited = true
}

// GetDB 获取db
func GetDB() *sql.DB {
	return mysqlDB
}
