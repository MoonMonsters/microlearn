/*
 * @Description:
 * @Author: your name
 * @Date: 2019-08-14 15:56:14
 * @LastEditTime: 2019-08-16 13:28:57
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	// 排序
	sort.Ints(nums)
	fmt.Println(nums)
	// 创建一个二维数组
	ret := make([][]int, 0, 0)
	for i := 0; i < len(nums); i = i + 1 {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 左右指针指向数组的两端
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					fmt.Println(nums[l], " ", nums[l+1])
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--

			} else if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			}
		}
	}
	return ret
}

func main() {
	nums := make([]int, 6)
	nums[0] = -2
	nums[1] = -1
	nums[2] = -2
	nums[3] = -2
	nums[4] = 2
	nums[5] = 4

	fmt.Println(threeSum(nums))
}
