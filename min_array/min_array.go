/*
 * @Author: linhuaqing
 * @Date: 2021-08-21 18:26:23
 * @LastEditTime: 2021-08-21 19:10:42
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/min_array/min_array.go
 * stay hungry stay foolish
 */
package min_array

// 暴力遍历，时间复杂度O(n)
func minArray1(numbers []int) int {
	length := len(numbers)
	index := 0
	for index < length-1 {
		if numbers[index] > numbers[index+1] {
			return numbers[index+1]
		}
		index++
	}
	return numbers[0]
}

// 二分法，时间复杂度O(logn)
// 关键在于
func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		mid := left + (right-left)/2
		if numbers[mid] < numbers[right] {
			right = mid
			continue
		}
		if numbers[mid] > numbers[right] {
			left = mid + 1
			continue
		}
		if numbers[mid] == numbers[right] {
			right--
			continue
		}
	}
	return numbers[left]
}
