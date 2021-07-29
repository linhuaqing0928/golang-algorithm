/*
 * @Author: linhuaqing
 * @Date: 2021-07-29 20:03:27
 * @LastEditTime: 2021-07-29 22:00:45
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/merge_two_lists/merge_two_lists.go
 * stay hungry stay foolish
 */
package merge_two_lists

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var result, current *ListNode
	if l1.Val <= l2.Val {
		current = l1
		l1 = l1.Next
	} else {
		current = l2
		l2 = l2.Next
	}
	result = current
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			current = current.Next
			l1 = l1.Next
			continue
		} else {
			current.Next = l2
			current = current.Next
			l2 = l2.Next
			continue
		}
	}
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}
	return result
}
