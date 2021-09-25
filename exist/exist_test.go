/*
 * @Author: linhuaqing
 * @Date: 2021-09-12 17:32:47
 * @LastEditTime: 2021-09-12 23:20:36
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/exist/exist_test.go
 * stay hungry stay foolish
 */
package exist

import "testing"

func TestExist(t *testing.T) {
	var board [][]byte
	var word string
	var result bool

	// board = [][]byte{
	// 	[]byte{'a', 'a'},
	// }
	// word = "a"
	// result = exist(board, word)
	// if result != true {
	// 	t.Errorf("Exist failed: %v", result)
	// }

	// board = [][]byte{
	// 	[]byte{'A', 'B', 'C', 'E'},
	// 	[]byte{'S', 'F', 'C', 'S'},
	// 	[]byte{'A', 'D', 'E', 'E'},
	// }
	// word = "ABCCED"
	// result = exist(board, word)
	// if result != true {
	// 	t.Errorf("Exist failed: %v", result)
	// }

	// board = [][]byte{
	// 	[]byte{'a', 'b'},
	// 	[]byte{'c', 'd'},
	// }
	// word = "abcd"
	// result = exist(board, word)
	// if result != false {
	// 	t.Errorf("Exist failed: %v", result)
	// }

	// board = [][]byte{
	// 	[]byte{'A', 'B', 'C', 'E'},
	// 	[]byte{'S', 'F', 'C', 'S'},
	// 	[]byte{'A', 'D', 'E', 'E'},
	// }
	// word = "SEE"
	// result = exist(board, word)
	// if result != true {
	// 	t.Errorf("Exist failed: %v", result)
	// }

	// board = [][]byte{
	// 	[]byte{'A', 'B', 'C', 'E'},
	// 	[]byte{'S', 'F', 'E', 'S'},
	// 	[]byte{'A', 'D', 'E', 'E'},
	// }
	// word = "ABCEFSADEESE"
	// result = exist(board, word)
	// if result != true {
	// 	t.Errorf("Exist failed: %v", result)
	// }

	// board = [][]byte{
	// 	[]byte{'a'},
	// }
	// word = "ab"
	// result = exist(board, word)
	// if result != false {
	// 	t.Errorf("Exist failed: %v", result)
	// }

	// board = [][]byte{
	// 	[]byte{'a', 'b'},
	// }
	// word = "ba"
	// result = exist(board, word)
	// if result != true {
	// 	t.Errorf("Exist failed: %v", result)
	// }

	// board = [][]byte{
	// 	[]byte{'a', 'a', 'b', 'a', 'a', 'b'},
	// 	[]byte{'a', 'a', 'b', 'b', 'b', 'a'},
	// 	[]byte{'a', 'a', 'a', 'a', 'b', 'a'},
	// 	[]byte{'b', 'a', 'b', 'b', 'a', 'b'},
	// 	[]byte{'a', 'b', 'b', 'a', 'b', 'a'},
	// 	[]byte{'b', 'a', 'a', 'a', 'a', 'b'},
	// }
	// word = "bbbaabbbbbab"
	// result = exist(board, word)
	// if result != false {
	// 	t.Errorf("Exist failed: %v", result)
	// }

	board = [][]byte{
		{'b', 'a', 'a', 'b', 'a', 'b'},
		{'a', 'b', 'a', 'a', 'a', 'a'},
		{'a', 'b', 'a', 'a', 'a', 'b'},
		{'a', 'b', 'a', 'b', 'b', 'a'},
		{'a', 'a', 'b', 'b', 'a', 'b'},
		{'a', 'a', 'b', 'b', 'b', 'a'},
		{'a', 'a', 'b', 'a', 'a', 'b'},
	}
	word = "aabbbbabbaababaaaabababbaaba"
	result = exist(board, word)
	if result != true {
		t.Errorf("Exist failed: %v", result)
	}

}
