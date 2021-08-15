/*
 * @Author: linhuaqing
 * @Date: 2021-08-15 21:07:27
 * @LastEditTime: 2021-08-16 01:13:27
 * @LastEditors: linhuaqing@bytedance.com
 * @Description:
 * @FilePath: /golang-algorithm/basic_calculate/basic_calculate_test.go
 * stay hungry stay foolish
 */
package basic_calculate

import (
	"fmt"
	"testing"
)

func TestCalCulate(t *testing.T) {
	var input string
	var result int
	var expected int

	input = "1 + 1"
	result = calculate(input)
	expected = 2
	fmt.Println(result)
	if result != expected {
		t.Errorf("TestcalCulate failed! expected %d, got %d", expected, result)
	}

	input = " 2-1 + 2 "
	result = calculate(input)
	expected = 3
	fmt.Println(result)
	if result != expected {
		t.Errorf("TestcalCulate failed! expected %d, got %d", expected, result)
	}

	input = "(1+(4+5+2)-3)+(6+8)"
	result = calculate(input)
	expected = 23
	fmt.Println(result)
	if result != expected {
		t.Errorf("TestcalCulate failed! expected %d, got %d", expected, result)
	}

	input = "2147483647"
	result = calculate(input)
	expected = 2147483647
	fmt.Println(result)
	if result != expected {
		t.Errorf("TestcalCulate failed! expected %d, got %d", expected, result)
	}

	input = "-2+ 1"
	result = calculate(input)
	expected = -1
	fmt.Println(result)
	if result != expected {
		t.Errorf("TestcalCulate failed! expected %d, got %d", expected, result)
	}

	input = "-(3+(4+5))"
	result = calculate(input)
	expected = -12
	fmt.Println(result)
	if result != expected {
		t.Errorf("TestcalCulate failed! expected %d, got %d", expected, result)
	}

	input = "(7)-(0)+(4)"
	result = calculate(input)
	expected = 11
	fmt.Println(result)
	if result != expected {
		t.Errorf("TestcalCulate failed! expected %d, got %d", expected, result)
	}
}
