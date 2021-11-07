/*
 * @Author: linhuaqing
 * @Date: 2021-10-24 17:52:53
 * @LastEditTime: 2021-10-31 23:12:16
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/is_symmetric/is_symmetric.go
 * stay hungry stay foolish
 */
package isSymmetric

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric([]*TreeNode{root.Left, root.Right})
}

func symmetric(roots []*TreeNode) bool {
	length := len(roots)
	index := 0
	for index < length/2 {
		temp1 := roots[index]
		temp2 := roots[length-1-index]
		if temp1 != nil && temp2 != nil && temp1.Val != temp2.Val {
			return false
		}
		if (temp1 != nil && temp2 == nil) || (temp1 == nil && temp2 != nil) {
			return false
		}
		index++
	}
	empty := true
	nextRoots := []*TreeNode{}
	index = 0
	for index < length {
		if roots[index] == nil {
			nextRoots = append(nextRoots, nil)
			nextRoots = append(nextRoots, nil)
			index++
			continue
		}
		nextRoots = append(nextRoots, roots[index].Left)
		if roots[index].Left != nil {
			empty = false
		}
		nextRoots = append(nextRoots, roots[index].Right)
		if roots[index].Right != nil {
			empty = false
		}
		index++
	}
	if !empty {
		return symmetric(nextRoots)
	}
	return true
}
