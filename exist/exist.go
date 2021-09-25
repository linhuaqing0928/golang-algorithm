/*
 * @Author: linhuaqing
 * @Date: 2021-09-12 16:14:03
 * @LastEditTime: 2021-09-13 01:10:48
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/exist/exist.go
 * stay hungry stay foolish
 */
package exist

import "fmt"

// func exist(board [][]byte, word string) bool {
// 	if len(word) > len(board)*len(board[0]) {
// 		return false
// 	}
// 	covers := new(list.List)
// 	index := 0
// 	outrow := 0
// 	outcolumn := 0

// 	for {
// 		if board[outrow][outcolumn] == word[0] {
// 			paths := make(map[[2]int]int)
// 			if outrow == 6 && outcolumn == 4 {
// 				fmt.Println("aaa")
// 			}
// 			fmt.Printf("start:row:%d,column:%d,value:%v\n", outrow, outcolumn, string(board[outrow][outcolumn]))
// 			// fmt.Println(string(board[outrow][outcolumn]))
// 			firstPath := [2]int{outrow, outcolumn}
// 			paths[firstPath] = 1
// 			covers.PushBack([]int{outrow - 1, outcolumn, 1, 1})
// 			covers.PushBack([]int{outrow, outcolumn - 1, 1, 2})
// 			covers.PushBack([]int{outrow + 1, outcolumn, 1, 3})
// 			covers.PushBack([]int{outrow, outcolumn + 1, 1, 4})
// 			fmt.Printf("covers size: %d\n", covers.Len())
// 			for covers.Len() != 0 && index < len(word) {
// 				removeNode := covers.Back()
// 				removeValue := covers.Remove(removeNode).([]int)
// 				row := removeValue[0]
// 				column := removeValue[1]
// 				index = removeValue[2]
// 				count := removeValue[3]
// 				if goOrnNot(board, word, row, column, index, paths) {
// 					fmt.Printf("row:%d, column:%d, index:%d,target:%v,current cover size: %d\n", row, column, index, string(word[index]), covers.Len())
// 					paths[[2]int{row, column}] = 1
// 					index++
// 					covers.PushBack([]int{row - 1, column, index, 1})
// 					fmt.Printf("push back row:%d,column:%d\n", row-1, column)
// 					covers.PushBack([]int{row, column - 1, index, 2})
// 					fmt.Printf("push back row:%d,column:%d\n", row, column-1)
// 					covers.PushBack([]int{row + 1, column, index, 3})
// 					fmt.Printf("push back row:%d,column:%d\n", row+1, column)
// 					covers.PushBack([]int{row, column + 1, index, 4})
// 					fmt.Printf("push back row:%d,column:%d\n", row, column+1)
// 					fmt.Printf("after push covers size:%d\n", covers.Len())
// 					continue
// 				} else {
// 					fmt.Printf("remove but not right:row:%d,column:%d\n", row, column)
// 					if count == 1 {
// 						fmt.Printf("revert!left stack size:%d\n", covers.Len())
// 						paths[[2]int{row + 1, column}] = 0
// 					}
// 				}
// 			}
// 			if index >= len(word)-1 {
// 				return true
// 			}
// 		}
// 		if outcolumn < len(board[outrow])-1 {
// 			outcolumn++
// 			index = 0
// 			continue
// 		}
// 		if outrow < len(board)-1 {
// 			outcolumn = 0
// 			outrow++
// 			index = 0
// 			continue
// 		}
// 		return false
// 	}
// }

// func goOrnNot(board [][]byte, word string, row int, column int, index int, paths map[[2]int]int) bool {
// 	if row < 0 || row >= len(board) || column < 0 || column >= len((board)[row]) {
// 		fmt.Println("row和column越界")
// 		return false
// 	}
// 	if index > len(word)-1 {
// 		fmt.Println("word index 越界")
// 		return false
// 	}
// 	if (board)[row][column] != word[index] {
// 		fmt.Println("board值不符合")
// 		return false
// 	}
// 	value, ok := paths[[2]int{row, column}]
// 	if ok && value == 1 {
// 		fmt.Println("已经经过")
// 		return false
// 	}
// 	// fmt.Println(string(word[index]))
// 	// fmt.Printf("row:%d, column:%d\n", row, column)
// 	return true
// }

func exist(board [][]byte, word string) bool {
	row := 0
	for row < len(board) {
		column := 0
		for column < len(board[0]) {
			if dfs(row, column, 0, &board, word) {
				// for _, rowTemp := range board {
				// 	rowTemp := rowTemp
				// 	for _, columnTemp := range rowTemp {
				// 		columnTemp := columnTemp
				// 		fmt.Printf("%v ", string(columnTemp))
				// 	}
				// 	fmt.Printf("\n")
				// }
				return true
			}
			column++
		}
		row++
	}
	return false
}

func dfs(i int, j int, k int, board *[][]byte, word string) bool {
	if !(0 <= i && i < len(*board)) || !(0 <= j && j < len((*board)[0])) || ((*board)[i][j] != word[k]) {
		return false
	}
	if k == len(word)-1 {
		fmt.Printf("row:%d,column:%d,value:%d\n", i, j, k)
		return true
	}
	(*board)[i][j] = ' '
	res := (dfs(i+1, j, k+1, board, word) || dfs(i-1, j, k+1, board, word) || dfs(i, j+1, k+1, board, word) || dfs(i, j-1, k+1, board, word))
	(*board)[i][j] = word[k]
	return res
}
