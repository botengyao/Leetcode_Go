/*
Package array ...
*/
package array

import (
	"sort"
)

//CombinationSum2 ...
func CombinationSum2(candidates []int, target int) [][]int {
	//回溯
	var temp []int
	var res [][]int
	sort.Ints(candidates)
	_backtrack(candidates, temp, 0, 0, target, &res)
	return res
}

func _backtrack(candidates []int, temp []int, start, total int, target int, res *[][]int) {
	if total == target {
		*res = append(*res, append([]int{}, temp...))
	} else if total < target {
		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			if total > target {
				break
			}
			backtrack(candidates, append(temp, candidates[i]), i+1, total+candidates[i], target, res)
		}
	}
}
