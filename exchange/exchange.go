/*
 * @Author: linhuaqing
 * @Date: 2021-09-04 11:54:44
 * @LastEditTime: 2021-09-04 12:15:58
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/exchange/exchange.go
 * stay hungry stay foolish
 */
package exchange

// 首位双指针
func exchange1(nums []int) []int {
	start := 0
	end := len(nums) - 1
	for start < end {
		if nums[start]%2 == 1 {
			start++
			continue
		}
		if nums[start]%2 == 0 {
			temp := nums[start]
			nums[start] = nums[end]
			nums[end] = temp
			end--
			continue
		}
	}
	return nums
}

// 快慢双指针
func exchange(nums []int) []int {
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast]%2 == 1 {
			temp := nums[fast]
			nums[fast] = nums[slow]
			nums[slow] = temp
			fast++
			slow++
			continue
		}
		if nums[fast]%2 == 0 {
			fast++
			continue
		}
	}
	return nums
}
