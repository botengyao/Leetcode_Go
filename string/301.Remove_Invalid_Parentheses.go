/*
Package string ...
*/
package string

func removeInvalidParentheses(s string) []string {
	//find the total mininum number
	//DFS for the string, remove or not remove 2^N
	//memorization
	//use a count to remember the number of deleted one
	//O(N * 2^N)
	res := make([]string, 0)
	visited := make(map[string]int)
	min := checkMinimum([]rune(s))
	removeInvalid(s, min, &res, &visited)
	return res
}

func removeInvalid(s string, min int, res *[]string, visited *map[string]int) {
	str := []rune(s)
	minnum := checkMinimum(str)
	(*visited)[s] = minnum
	if minnum > min {
		return
	}
	if minnum == 0 {
		*res = append(*res, s)
	}

	for i := 0; i < len(str); i++ {
		if str[i] != '(' && str[i] != ')' {
			continue
		}
		newStr := make([]rune, 0)
		newStr = append(newStr, str[:i]...)
		newStr = append(newStr, str[i+1:]...)
		temp := string(newStr)
		if _, ok := (*visited)[temp]; !ok {
			removeInvalid(temp, min-1, res, visited)
		}
	}
}

func checkMinimum(s []rune) int {
	stack := make([]rune, 0)
	res := 0
	for _, char := range s {
		if char == '(' {
			stack = append(stack, char)
		} else if char == ')' {
			if len(stack) > 0 {
				stack = stack[0 : len(stack)-1]
			} else {
				res++
			}
		}
	}
	return res + len(stack)
}
