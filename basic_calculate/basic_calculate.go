/*
 * @Author: linhuaqing
 * @Date: 2021-08-15 17:54:02
 * @LastEditTime: 2021-08-16 01:16:24
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/basic_calculate/basic_calculate.go
 * stay hungry stay foolish
 */
package basic_calculate

import (
	"container/list"
	"strconv"
	"strings"
)

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	var result int
	intStack := list.New()
	intStack.PushBack(0)
	oprateStack := list.New()
	start := 0
	end := start + 1
	for end <= len(s) {
		if s[start:end] == "(" {
			oprateStack.PushBack("(")
			start = end
			end = start + 1
			continue
		}
		if s[start:end] == "-" || s[start:end] == "+" {
			if start > 0 && (s[start-1:start] == "-" || s[start-1:start] == "+" || s[start-1:start] == "(") {
				intStack.PushBack(0)
			}
			for oprateStack.Len() != 0 && oprateStack.Back().Value != "(" {
				topInt := intStack.Back()
				intLast := intStack.Remove(topInt).(int)
				topInt1 := intStack.Back()
				intLast1 := intStack.Remove(topInt1).(int)
				topOprate := oprateStack.Back()
				oprateLast := oprateStack.Remove(topOprate).(string)
				tempResult := caculateInnner(intLast1, intLast, oprateLast)
				intStack.PushBack(tempResult)
			}
			oprateStack.PushBack(s[start:end])
			start = end
			end = start + 1
			continue
		}
		if s[start:end] == ")" {
			for oprateStack.Len() != 0 && oprateStack.Back().Value != "(" {
				topInt := intStack.Back()
				intLast := intStack.Remove(topInt).(int)
				topInt1 := intStack.Back()
				intLast1 := intStack.Remove(topInt1).(int)
				topOprate := oprateStack.Back()
				oprateLast := oprateStack.Remove(topOprate).(string)
				tempResult := caculateInnner(intLast1, intLast, oprateLast)
				intStack.PushBack(tempResult)
			}
			if oprateStack.Len() != 0 && oprateStack.Back().Value == "(" {
				topOprate := oprateStack.Back()
				oprateStack.Remove(topOprate)
			}
			start = end
			end = start + 1
			continue
		}
		if end+1 <= len(s) && s[end:end+1] != "+" && s[end:end+1] != "-" && s[end:end+1] != "(" && s[end:end+1] != ")" {
			end += 1
			continue
		}
		tempInt, _ := strconv.Atoi(s[start:end])
		intStack.PushBack(tempInt)
		start = end
		end = start + 1
	}
	for oprateStack.Len() != 0 {
		topInt := intStack.Back()
		intLast := intStack.Remove(topInt).(int)
		topInt1 := intStack.Back()
		intLast1 := intStack.Remove(topInt1).(int)
		topOprate := oprateStack.Back()
		oprateLast := oprateStack.Remove(topOprate).(string)
		tempResult := caculateInnner(intLast1, intLast, oprateLast)
		intStack.PushBack(tempResult)
	}
	topInt := intStack.Back()
	result = intStack.Remove(topInt).(int)
	return result
}

func caculateInnner(a, b int, sumbol string) int {
	if sumbol == "+" {
		return a + b
	} else {
		return a - b
	}
}
