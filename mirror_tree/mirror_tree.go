/*
 * @Author: linhuaqing
 * @Date: 2021-08-04 22:21:39
 * @LastEditTime: 2021-08-08 12:34:09
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/mirror_tree/mirror_tree.go
 * stay hungry stay foolish
 */
package mirror_tree

import "container/list"

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归算法
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	if root.Right != nil {
		mirrorTree(root.Right)
	}
	if root.Left != nil {
		mirrorTree(root.Left)
	}
	return root
}

// 使用栈解法
func mirrorTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	list := list.New()
	list.PushBack(root)
	for list.Len() != 0 {
		first := list.Back()
		firstEle := list.Remove(first).(*TreeNode)
		if firstEle.Left != nil || firstEle.Right != nil {
			temp := firstEle.Left
			firstEle.Left = firstEle.Right
			firstEle.Right = temp
		}
		if firstEle.Left != nil {
			list.PushBack(firstEle.Left)
		}
		if firstEle.Right != nil {
			list.PushBack(firstEle.Right)
		}
	}
	return root
}
