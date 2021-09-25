/*
 * @Author: linhuaqing
 * @Date: 2021-09-25 22:40:22
 * @LastEditTime: 2021-09-25 22:44:14
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/spiral_order/spiral_order_test.go
 * stay hungry stay foolish
 */
package spiral_order

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralOrder(t *testing.T) {
	var matrix [][]int
	var result []int
	var expected []int

	matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected = []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	result = spiralOrder(matrix)
	assert.Equal(t, expected, result, "they should be equal")
}
