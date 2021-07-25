/*
 * @Author: linhuaqing
 * @Date: 2021-07-21 21:53:23
 * @LastEditTime: 2021-07-22 10:08:54
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/build_tree/build_tree.go
 * stay hungry stay foolish
 */
package build_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 遍历中序数组并将数据放到map中，方便后续读取数据
	index := 0
	treeMap := make(map[int]int)
	for ; index < len(inorder); index++ {
		treeMap[inorder[index]] = index
	}
	head := buildTreeWithIndex(0, len(preorder)-1, 0, len(preorder)-1, preorder, inorder, treeMap)
	return head
}

func buildTreeWithIndex(preorderLeft, preorderRight, inorderLeft, inorderRight int, preorder []int, inorder []int, treeMap map[int]int) *TreeNode {
	if preorderLeft > preorderRight {
		return nil
	}

	preorderRoot := preorderLeft
	inorderRoot := treeMap[preorder[preorderRoot]]
	inorderLeftSize := inorderRoot - inorderLeft
	head := &TreeNode{preorder[preorderRoot], nil, nil}
	head.Left = buildTreeWithIndex(preorderLeft+1, preorderLeft+inorderLeftSize, inorderLeft, inorderRoot-1, preorder, inorder, treeMap)
	head.Right = buildTreeWithIndex(preorderLeft+inorderLeftSize+1, preorderRight, inorderRoot+1, inorderRight, preorder, inorder, treeMap)
	return head
}
