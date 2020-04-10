/*
Package string ...
*/
package string

import (
	"strings"
)

/*
Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.

Example 1:

Input: S = "ab#c", T = "ad#c"
Output: true
Explanation: Both S and T become "ac".
Example 2:

Input: S = "ab##", T = "c#d#"
Output: true
Explanation: Both S and T become "".
Example 3:

Input: S = "a##c", T = "#a#c"
Output: true
Explanation: Both S and T become "c".
Example 4:

Input: S = "a#c", T = "b"
Output: false
Explanation: S becomes "c" while T becomes "b".
*/

//BackspaceCompare from end to start
func BackspaceCompare(S string, T string) bool {
	//stack O(N)
	//pointer O(1)
	//from last to start, once encounter a "#", count++ and keep move i--. if the remaining "#" is bigger than 0 and the i position is not “#”, count--, i--
	//that means we ignore the position. Then compare s[i] and t[j]
	s := strings.Split(S, "")
	t := strings.Split(T, "")
	i, j := len(s)-1, len(t)-1
	countT, countS := 0, 0
	for {
		for ; i >= 0 && (countS > 0 || s[i] == "#"); i-- {
			if s[i] == "#" {
				countS++
			} else {
				countS--
			}
		}

		for ; j >= 0 && (countT > 0 || t[j] == "#"); j-- {
			if t[j] == "#" {
				countT++
			} else {
				countT--
			}
		}

		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
			i--
			j--

		} else {
			return i < 0 && j < 0
		}
	}
}
