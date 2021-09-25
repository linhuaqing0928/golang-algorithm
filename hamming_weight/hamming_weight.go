/*
 * @Author: linhuaqing
 * @Date: 2021-09-21 08:44:52
 * @LastEditTime: 2021-09-21 09:05:50
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/hamming_weight/hamming_weight.go
 * stay hungry stay foolish
 */
package hamming_weight

import "math"

func hammingWeight1(num uint32) int {
	count := 0
	index := 31
	for index >= 0 {
		multi := num / (uint32(math.Pow(2, float64(index))))
		if multi == 1 {
			count++
		}
		num = num % (uint32(math.Pow(2, float64(index))))
		index--
	}
	return count
}

func hammingWeight(num uint32) int {
	count := 0
	index := 0
	for index < 32 {
		if 1<<index&num > 0 {
			count++
		}
		index++
	}
	return count
}
