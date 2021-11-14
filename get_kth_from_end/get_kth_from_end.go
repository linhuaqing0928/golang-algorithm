/*
 * @Author: linhuaqing
 * @Date: 2021-11-14 17:25:44
 * @LastEditTime: 2021-11-14 18:06:48
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/get_kth_from_end/get_kth_from_end.go
 * stay hungry stay foolish
 */
package get_kth_from_end

type ListNode struct {
	Val  int
	Next *ListNode
}

// map存储法
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	nodes := make(map[int]*ListNode)
	count := 1
	for {
		nodes[count] = head
		if head.Next != nil {
			head = head.Next
			count++
		} else {
			break
		}
	}
	result := nodes[count-k+1]
	return result
}

// 快慢指针方法
func getKthFromEnd1(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	index := 1
	for {
		if fast.Next != nil {
			if index < k {
				fast = fast.Next
				index++
			} else {
				fast = fast.Next
				slow = slow.Next
			}
		} else {
			break
		}
	}
	return slow
}
