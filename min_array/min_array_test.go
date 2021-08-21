/*
 * @Author: linhuaqing
 * @Date: 2021-08-21 18:33:56
 * @LastEditTime: 2021-08-21 19:09:53
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/min_array/min_array_test.go
 * stay hungry stay foolish
 */
package min_array

import "testing"

func TestMinArray(t *testing.T) {
	var input []int
	var result int
	var expected int

	input = []int{3, 4, 5, 1, 2}
	result = minArray(input)
	expected = 1
	if result != expected {
		t.Errorf("TestMinArray failed,should be %d, got %d", expected, result)
	}

	input = []int{2, 2, 2, 0, 1}
	result = minArray(input)
	expected = 0
	if result != expected {
		t.Errorf("TestMinArray failed,should be %d, got %d", expected, result)
	}

	input = []int{3, 1, 3}
	result = minArray(input)
	expected = 1
	if result != expected {
		t.Errorf("TestMinArray failed,should be %d, got %d", expected, result)
	}
}
