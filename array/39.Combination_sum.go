/*
Package array ...
*/
package array

import (
	"sort"
)

//CombinationSum ...
func CombinationSum(candidates []int, target int) [][]int {
	//回溯
	var temp []int
	var res [][]int
	sort.Ints(candidates)
	backtrack(candidates, temp, 0, 0, target, &res)
	return res
}

func backtrack(candidates []int, temp []int, start, total int, target int, res *[][]int) {
	if total == target {
		*res = append(*res, append([]int{}, temp...))
	} else if total < target {
		for i := start; i < len(candidates); i++ {
			backtrack(candidates, append(temp, candidates[i]), i, total+candidates[i], target, res)
		}
	}
}
