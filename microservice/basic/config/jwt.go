/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-18 08:10:42
 * @LastEditTime: 2019-08-18 08:13:17
 * @LastEditors: Please set LastEditors
 */
package config

type JwtConfig interface {
	GetSecretKey() string
}
type defaultJwtConfig struct {
	SecretKey string `json:"secretKey"`
}

func (m defaultJwtConfig) GetSecretKey() string {
	return m.SecretKey
}
