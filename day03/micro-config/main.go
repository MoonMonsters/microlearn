/*
 * @Description: In User Settings Edit
 * @Author: 陌无崖
 * @Date: 2019-08-15 09:13:26
 * @LastEditTime: 2019-08-15 14:05:05
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
)

func main() {
	// 加载配置文件
	if err := config.Load(file.NewSource(
		file.WithPath("./config/config.json"),
	// file.WithPath("./config/config.yml"),
	)); err != nil {
		fmt.Println(err)
		return
	}

	// 定义我们的额数据结构
	type Host struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Port    int    `json:"port"`
	}
	var host Host

	if err := config.Get("hosts", "database").Scan(&host); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(host.Name, host.Address, host.Port)
	// 测试hosttwo
	var hosttwo Host

	if err := config.Get("hosttwo").Scan(&hosttwo); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hosttwo.Name, hosttwo.Address, hosttwo.Port)
	// 测试hostthree

	type HostAddress struct {
		Database `json:"database"`
	}

	var hostthree HostAddress
	if err := config.Get("hostthree").Scan(&hostthree); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hostthree.Database)
	// 测试hostfour
	type HostFour struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Port    int    `json:"port"`
	}
	var hostfour []HostFour
	if err := config.Get("hostfour").Scan(&hostfour); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hostfour)
}

type Database struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}
