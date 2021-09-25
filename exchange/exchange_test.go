/*
 * @Author: linhuaqing
 * @Date: 2021-09-04 11:58:14
 * @LastEditTime: 2021-09-04 12:05:22
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/exchange/exchange_test.go
 * stay hungry stay foolish
 */
package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExchange(t *testing.T) {
	var input []int
	var expected []int
	var result []int

	input = []int{1, 2, 3, 4}
	expected = []int{1, 3, 4, 2}
	result = exchange(input)
	assert.Equal(t, expected, result, "they should be equal")
}
