/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-19 09:15:35
 * @LastEditTime: 2019-08-19 09:22:23
 * @LastEditors: Please set LastEditors
 */
 package inventory

 import (
	 "fmt"
	 "sync"
 )
 
 var (
	 s *service
	 m sync.RWMutex
 )
 
 type Service interface {
	 Sell(bookId int64, userId int64) (id int64, err error)
	 // Confirm 确认销存
	 Confirm(id int64, state int) (err error)
 }
 type service struct {
 }
 
 func GetService() (Service, error) {
	 if s == nil {
		 return nil, fmt.Errorf("[GetService] GetService 未初始化")
	 }
	 return s, nil
 }
 
 // Init 初始化库存服务层
 func Init() {
	 m.Lock()
	 defer m.Unlock()
 
	 if s != nil {
		 return
	 }
 
	 s = &service{}
 }
 