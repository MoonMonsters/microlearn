/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-16 17:37:41
 * @LastEditTime: 2019-08-16 17:42:01
 * @LastEditors: Please set LastEditors
 */
package config

// 定义MySql配置的接口
type MysqlConfig interface {
	GetURL() string
	GetEnabled() bool
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
}

// defaultMysqlConfig mysql 配置
type defaultMysqlConfig struct {
	URL               string `json:"url"`
	Enable            bool   `json:"enabled"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
}

// URL mysql 连接
func (m defaultMysqlConfig) GetURL() string {
	return m.URL
}

// Enabled 激活
func (m defaultMysqlConfig) GetEnabled() bool {
	return m.Enable
}

// 闲置连接数
func (m defaultMysqlConfig) GetMaxIdleConnection() int {
	return m.MaxIdleConnection
}

// 打开连接数
func (m defaultMysqlConfig) GetMaxOpenConnection() int {
	return m.MaxOpenConnection
}
