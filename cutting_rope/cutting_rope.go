/*
 * @Author: linhuaqing
 * @Date: 2021-10-03 17:20:04
 * @LastEditTime: 2021-10-07 18:07:24
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/cutting_rope/cutting_rope.go
 * stay hungry stay foolish
 */
package cutting_rope

import (
	"math"
)

// 贪心算法 找规律，可以得到将内容划分为3的时候 得到的乘积为最大
func cuttingRope(n int) int {
	var result int
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	remind := n % 3
	if remind == 0 {
		// result = int(math.Pow(3, float64(n/3))) % 1000000007
		result = getRemind(n/3) % 1000000007
	}
	if remind == 1 {
		// result = int(math.Pow(3, float64(n/3)-1)) * 4 % 1000000007
		result = getRemind(n/3-1) * 4 % 1000000007
	}
	if remind == 2 {
		// result = int(math.Pow(3, float64(n/3))) * 2 % 1000000007
		result = getRemind(n/3) * 2 % 1000000007
	}
	return result
}

func getRemind(n int) int {
	index := 1
	var result int
	result = 1
	for index <= n {
		result = result * 3 % 1000000007
		index++
	}
	return result
}

// 动态规划
// 因为必须m>1，所以必须剪一次 假设剪掉的一段长度为J，则剩余的长度为n-j 剩余的绳子可以有2种选择：剪或者不剪 剪的话，就按照之前最优解的乘积*j 不剪的话，就是（n-j）*j 取max得出最优解
// J的长度是变化的，但是至少从2开始 这里也要取max
func cuttingRope1(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	i := 3
	for i <= n {
		j := 2
		for j < i {
			dp[i] = int(math.Max(float64(dp[i]), math.Max(float64(j*(i-j)), float64(j*(dp[i-j])))))
			j++
		}
		i++
	}
	return dp[n]
}
