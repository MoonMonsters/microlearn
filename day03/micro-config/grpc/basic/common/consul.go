/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-24 14:08:39
 * @LastEditTime: 2019-08-24 14:08:47
 * @LastEditors: Please set LastEditors
 */
package common

// Consul 配置
type Consul struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}