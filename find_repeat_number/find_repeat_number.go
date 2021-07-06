/*
 * @Author: linhuaqing
 * @Date: 2021-07-06 20:24:36
 * @LastEditTime: 2021-07-06 20:59:40
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/find_repeat_number/find_repeat_number.go
 * stay hungry stay foolish
 */
package find_repeat_number

// map不可重复解法
func findRepeatNumber(nums []int) int {
	unrepeat_map := make(map[int]int)
	for _, num := range nums {
		if _, ok := unrepeat_map[num]; ok {
			return num
		} else {
			unrepeat_map[num] = 1
		}
	}
	return -1
}

// 原地交换算法
func findRepeatNumber1(nums []int) int {
	index := 0
	for index < len(nums) {
		if index == nums[index] {
			index++
			continue
		}
		if nums[nums[index]] == nums[index] {
			return nums[index]
		}
		temp := nums[nums[index]]
		nums[nums[index]] = nums[index]
		nums[index] = temp
	}
	return -1
}
