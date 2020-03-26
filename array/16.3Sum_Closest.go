/*
Package array ...
*/
package array

import (
	"math"
	"sort"
)

//ThreeSumClosest ...
func ThreeSumClosest(nums []int, target int) int {
	//sort
	//和当前比较 如果大就减
	//如果小就更新
	sort.Ints(nums)
	min, left, right := math.MaxFloat64, 0, len(nums)-1
	res, pair, diff := 0, 0, 0
	for i := 0; i < len(nums)-2; i++ {
		left, right = i+1, len(nums)-1
		for left < right {
			pair = nums[left] + nums[right]
			diff = pair + nums[i] - target
			if math.Abs(float64(diff)) < min { //it is not a good way to change the data type....
				res = pair + nums[i]
				min = math.Abs(float64(diff))
			}
			if diff < 0 {
				left++
			} else if diff > 0 {
				right--
			} else {
				break
			}
		}
	}
	return res
}
