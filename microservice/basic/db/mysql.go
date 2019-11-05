/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-16 18:31:38
 * @LastEditTime: 2019-08-18 12:31:33
 * @LastEditors: Please set LastEditors
 */
package db

import (
	"database/sql"
	"sync"

	"github.com/micro/go-micro/util/log"
	"github.com/yuwe1/micolearn/microservice/basic/config"
)

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

func initMysql() {
	var err error

	// 创建连接
	mysqlDB, err = sql.Open("mysql", config.GetMysqlConfig().GetURL())

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	// 最大连接数
	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())

	// 最大闲置数
	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())

	// 激活链接
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
}
