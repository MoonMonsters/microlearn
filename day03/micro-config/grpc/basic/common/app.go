/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-24 14:08:14
 * @LastEditTime: 2019-08-24 14:08:21
 * @LastEditors: Please set LastEditors
 */
package common

import "strconv"

// AppCfg common config
type AppCfg struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

func (a *AppCfg) Addr() string {
	return a.Address + ":" + strconv.Itoa(a.Port)
}