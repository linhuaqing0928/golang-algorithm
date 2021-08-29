/*
 * @Author: linhuaqing
 * @Date: 2021-08-29 10:06:59
 * @LastEditTime: 2021-08-29 10:38:26
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/num_ways/num_ways.go
 * stay hungry stay foolish
 */
package num_ways

func numWays(n int) int {
	if n < 2 {
		return 1
	}
	result := make([]int, n+1)
	result[0] = 1
	result[1] = 1
	index := 2
	for index <= n {
		result[index] = result[index-1] + result[index-2]
		result[index] = result[index] % 1000000007
		index++
	}
	return result[n]
}
