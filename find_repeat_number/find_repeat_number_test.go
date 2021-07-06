/*
 * @Author: linhuaqing
 * @Date: 2021-07-06 20:33:07
 * @LastEditTime: 2021-07-06 20:59:00
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/find_repeat_number/find_repeat_number_test.go
 * stay hungry stay foolish
 */
package find_repeat_number

import (
	"fmt"
	"testing"
)

func TestFindRepeatNumber(t *testing.T) {
	var input []int
	var result int

	input = []int{2, 3, 1, 0, 2, 5, 3}
	result = findRepeatNumber(input)
	fmt.Printf("result1 : %d\n", result)
	if result != 2 && result != 3 {
		t.Errorf("用例1失败")
	}

}

func TestFindRepeatNumber1(t *testing.T) {
	var input []int
	var result int

	input = []int{2, 3, 1, 0, 2, 5, 3}
	result = findRepeatNumber1(input)
	fmt.Printf("result1 : %d\n", result)
	if result != 2 && result != 3 {
		t.Errorf("用例1失败")
	}

}
