/*
 * @Author: linhuaqing
 * @Date: 2021-08-29 10:11:44
 * @LastEditTime: 2021-08-29 10:38:49
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/num_ways/num_ways_test.go
 * stay hungry stay foolish
 */
package num_ways

import "testing"

func TestNumWays(t *testing.T) {
	var input int
	var result int
	var expected int

	input = 2
	expected = 2
	result = numWays(input)
	if result != expected {
		t.Errorf("should be %d, got %d", expected, result)
	}

	input = 7
	expected = 21
	result = numWays(input)
	if result != expected {
		t.Errorf("should be %d, got %d", expected, result)
	}

	input = 0
	expected = 1
	result = numWays(input)
	if result != expected {
		t.Errorf("should be %d, got %d", expected, result)
	}

	input = 100
	expected = 1
	result = numWays(input)
	if result != expected {
		t.Errorf("should be %d, got %d", expected, result)
	}

}
