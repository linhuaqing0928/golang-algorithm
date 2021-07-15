/*
 * @Author: linhuaqing
 * @Date: 2021-07-15 19:52:19
 * @LastEditTime: 2021-07-15 21:10:43
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/reverse_print/reverse_print.go
 * stay hungry stay foolish
 */
package reverse_print

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var next *ListNode
	var prev *ListNode
	for true {
		if head.Next != nil {
			next = head.Next
			head.Next = prev
			prev = head
			head = next
			continue
		}
		head.Next = prev
		break
	}
	var result []int
	for true {
		result = append(result, head.Val)
		if head.Next != nil {
			head = head.Next
			continue
		}
		break
	}
	return result
}
