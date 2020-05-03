/*
Package string ...
*/
package string

import (
	"strings"
)

//MinRemoveToMakeValid ...
func MinRemoveToMakeValid(s string) string {
	//左括号多了
	//右括号多了
	stack := make([]int, 0)
	str := []rune(s)

	for i, char := range str {
		if char == '(' {
			stack = append(stack, i)
		} else if char == ')' {
			if len(stack) > 0 {
				stack = stack[0 : len(stack)-1]
			} else {
				str[i] = '.'
			}
		}
	}

	for i := len(stack) - 1; i >= 0; i-- {
		str[stack[i]] = '.'
	}

	return strings.ReplaceAll(string(str), ".", "")
}
