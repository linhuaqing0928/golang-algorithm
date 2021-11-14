/*
 * @Author: linhuaqing
 * @Date: 2021-11-14 17:42:50
 * @LastEditTime: 2021-11-14 18:05:28
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/get_kth_from_end/get_kth_from_end_test.go
 * stay hungry stay foolish
 */
package get_kth_from_end

import "testing"

func TestGetKthFromEnd(t *testing.T) {
	var input *ListNode
	var result *ListNode
	var k int

	input = BuildTree([]int{1, 2, 3, 4, 5})
	k = 2
	result = getKthFromEnd1(input, k)
	if result.Val != 4 {
		t.Errorf("TestGetKthFromEnd failed, got %v,expected %v", result.Val, 4)
	}
}

func BuildTree(inputInts []int) *ListNode {
	head := &ListNode{inputInts[0], nil}
	result := head
	for index, inputInt := range inputInts {
		index := index
		inputInt := inputInt
		if index != 0 {
			newNode := &ListNode{inputInt, nil}
			head.Next = newNode
			head = head.Next
		}
	}
	return result
}
