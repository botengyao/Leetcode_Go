/*
Package string ...
*/
package string

/*
Given a string containing only three types of characters: '(', ')' and '*', write a function to check whether this string is valid. We define the validity of a string by these rules:

Any left parenthesis '(' must have a corresponding right parenthesis ')'.
Any right parenthesis ')' must have a corresponding left parenthesis '('.
Left parenthesis '(' must go before the corresponding right parenthesis ')'.
'*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string.
An empty string is also valid.
Example 1:
Input: "()"
Output: True
Example 2:
Input: "(*)"
Output: True
Example 3:
Input: "(*))"
Output: True
*/

//CheckValidString ...
func CheckValidString(s string) bool {
	// *(   *)
	// too many right "("   * or ")" count
	// too many right ")"
	leftwild := 0
	left := 0
	runeArray := []rune(s)
	for _, char := range runeArray {
		//右括号太多了
		if char == ')' {
			if leftwild == 0 {
				return false
			}
			leftwild--
		} else {
			leftwild++
		}

		//左括号太多了
		if char == '(' {
			left++
		} else {
			left = maxVal(0, left-1) // ((******) -> * can be empty
		}
	}
	return left == 0
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
