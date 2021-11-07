/*
 * @Author: linhuaqing
 * @Date: 2021-10-03 18:04:56
 * @LastEditTime: 2021-10-07 18:07:45
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/cutting_rope/cutting_rope_test.go
 * stay hungry stay foolish
 */
package cutting_rope

import "testing"

func TestCuttingRope(t *testing.T) {
	var input int
	var result int
	var expected int

	// input = 2
	// expected = 1
	// result = cuttingRope(input)
	// if result != expected {
	// 	t.Errorf("cuttingRope failed: expected %v, got %v", expected, result)
	// }

	input = 10
	expected = 36
	result = cuttingRope(input)
	if result != expected {
		t.Errorf("cuttingRope failed: expected %v, got %v", expected, result)
	}

	input = 120
	expected = 953271190
	result = cuttingRope(input)
	if result != expected {
		t.Errorf("cuttingRope failed: expected %v, got %v", expected, result)
	}
}
