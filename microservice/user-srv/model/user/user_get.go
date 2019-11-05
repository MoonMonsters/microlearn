/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-17 08:49:17
 * @LastEditTime: 2019-08-18 12:32:59
 * @LastEditors: Please set LastEditors
 */
package user

import (
	"github.com/micro/go-micro/util/log"
	"github.com/yuwe1/micolearn/microservice/basic/db"
	proto "github.com/yuwe1/micolearn/microservice/user-srv/proto/user"
)

func (s *service) QueryUserByName(userName string) (ret *proto.User, err error) {
	queryString := `SELECT user_id, user_name, pwd FROM user WHERE user_name = ?`
	// 获得数据库
	o := db.GetDB()

	ret = &proto.User{}

	// 查询
	err = o.QueryRow(queryString, userName).Scan(&ret.Id, &ret.Name, &ret.Pwd)
	if err != nil {
		log.Logf("[QueryUserByName] 查询数据失败，err：%s", err)
		return
	}
	return
}
