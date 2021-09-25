/*
 * @Author: linhuaqing
 * @Date: 2021-09-25 21:59:45
 * @LastEditTime: 2021-09-25 23:59:43
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/spiral_order/spiral_order.go
 * stay hungry stay foolish
 */
package spiral_order

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	result := make([]int, 0, len(matrix)*len(matrix[0]))
	xMin := 0
	yMin := 1 // 这里需要注意，Y轴刚起步就占用了一个，所以初始为1
	yMax := len(matrix)
	xMax := len(matrix[0])
	x := 0       // 代表X轴
	y := 0       // 代表Y轴
	flag := true // true代表横着走，false代表竖着走
	count := 1   // 1代表正向走，-1达标逆向走
	for {
		if flag {
			if x < xMin || x >= xMax {
				break
			}
		} else {
			if y < yMin || y >= yMax {
				break
			}
		}
		result = append(result, matrix[y][x])
		// 专项判断
		if flag {
			xTemp := x + count
			if xTemp >= xMax {
				flag = false
				count = 1
				xMax--
			} else if xTemp < xMin {
				flag = false
				count = -1
				xMin++
			}
		} else if !flag {
			yTemp := y + count
			if yTemp >= yMax {
				flag = true
				count = -1
				yMax--
			} else if yTemp < yMin {
				flag = true
				count = 1
				yMin++
			}
		}
		// 正式开始转向
		if flag {
			x += count
		} else {
			y += count
		}
	}
	return result
}
