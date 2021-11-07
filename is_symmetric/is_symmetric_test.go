/*
 * @Author: linhuaqing
 * @Date: 2021-10-24 18:15:43
 * @LastEditTime: 2021-10-31 23:14:20
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/is_symmetric/is_symmetric_test.go
 * stay hungry stay foolish
 */
package isSymmetric

import (
	"math"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	input := []int{1, 2, 2, 3, 4, 4, 3}
	root := buildTree(input)
	// fmt.Println(root)
	result1 := isSymmetric(root)
	if result1 != true {
		t.Errorf("TestIsSymmetric failed, got %v, want %v", result1, true)
	}
}

func buildTree(input []int) *TreeNode {
	root := &TreeNode{input[0], nil, nil}
	lastNodes := []*TreeNode{root}
	nextNodes := []*TreeNode{}
	level := 1
	start := 1
	end := start + int(math.Pow(2, float64(level)))
	for end <= len(input) {
		for index, node := range lastNodes {
			index := index
			node := node
			if node == nil {
				continue
			} else {
				first := input[start+2*index]
				if first != -1 {
					firstNode := &TreeNode{first, nil, nil}
					node.Left = firstNode
					nextNodes = append(nextNodes, firstNode)
				} else {
					nextNodes = append(nextNodes, nil)
				}
				second := input[start+2*index+1]
				if second != -1 {
					secondNode := &TreeNode{second, nil, nil}
					node.Right = secondNode
					nextNodes = append(nextNodes, secondNode)
				} else {
					nextNodes = append(nextNodes, nil)
				}
			}
		}
		lastNodes = nextNodes
		nextNodes = []*TreeNode{}
		start = end
		level++
		end = start + int(math.Pow(2, float64(level)))
	}
	return root
}
