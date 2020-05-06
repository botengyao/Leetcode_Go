/*
Package string ...
*/
package string

import (
	"reflect"
)

func findAnagrams(s string, p string) []int {
	//"abc" -> sort
	res := make([]int, 0)
	memo := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		memo[p[i]]++
	}

	slide := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		slide[s[i]]++
		if i >= len(p) {
			slide[s[i-len(p)]]--
			if slide[s[i-len(p)]] == 0 {
				delete(slide, s[i-len(p)])
			}
		}
		if reflect.DeepEqual(slide, memo) {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}
