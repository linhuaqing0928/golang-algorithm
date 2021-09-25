/*
 * @Author: linhuaqing
 * @Date: 2021-09-21 08:51:07
 * @LastEditTime: 2021-09-21 08:53:31
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/hamming_weight/hamming_weight_test.go
 * stay hungry stay foolish
 */
package hamming_weight

import "testing"

func TestHammingWeight(t *testing.T) {
	var input uint32
	var result int
	var expected int

	input = 11
	result = hammingWeight(input)
	expected = 3
	if result != expected {
		t.Errorf("hamming_weight failed: expected %v, got %v", expected, result)
	}

	input = 128
	result = hammingWeight(input)
	expected = 1
	if result != expected {
		t.Errorf("hamming_weight failed: expected %v, got %v", expected, result)
	}

	input = 4294967293
	result = hammingWeight(input)
	expected = 31
	if result != expected {
		t.Errorf("hamming_weight failed: expected %v, got %v", expected, result)
	}
}
