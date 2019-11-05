/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-16 17:49:03
 * @LastEditTime: 2019-08-16 18:04:55
 * @LastEditors: Please set LastEditors
 */
package config

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestInit(t *testing.T) {
	sp := "/"
	currentpath := filepath.Dir(filepath.Join("."+sp, sp))
	abs, _ := filepath.Abs(currentpath)
	pt := filepath.Join(abs, "../../conf")
	fmt.Println(pt)
}
