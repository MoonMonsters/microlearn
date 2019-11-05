/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-24 14:02:40
 * @LastEditTime: 2019-08-24 14:03:09
 * @LastEditors: Please set LastEditors
 */
package gconfig

import "github.com/micro/go-micro/config/source"

type Options struct {
	Apps    map[string]interface{}
	Sources []source.Source
}

type Option func(o *Options)

func WithSource(src source.Source) Option {
	return func(o *Options) {
		o.Sources = append(o.Sources, src)
	}
}