/*
 * @Author: linhuaqing
 * @Date: 2021-07-15 21:12:35
 * @LastEditTime: 2021-07-15 21:34:55
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/reverse_print/reverse_print_test.go
 * stay hungry stay foolish
 */
package reverse_print

import (
	"fmt"
	"testing"
)

func TestReversePrint(t *testing.T) {
	var input []int
	var head *ListNode
	var result []int
	fmt.Println("aaaaa")

	// example test
	input = []int{1, 3, 2}
	head = createListNodeFromList(input)
	result = reversePrint(head)
	fmt.Println(result)
	if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", []int{2, 3, 1}) {
		t.Errorf("example test failed!")
	}

	// example test
	input = []int{}
	head = createListNodeFromList(input)
	result = reversePrint(head)
	fmt.Println(result)
	if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", []int{}) {
		t.Errorf("example test failed!")
	}
}

func createListNodeFromList(input []int) *ListNode {
	var head *ListNode
	var prev *ListNode
	for i := 0; i < len(input); i++ {
		temp := &ListNode{input[i], nil}
		if prev == nil {
			prev = temp
			head = prev
		} else {
			prev.Next = temp
			prev = temp
		}
	}
	return head
}
