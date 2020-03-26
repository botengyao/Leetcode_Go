/*
Package array ...
*/
package array

import (
	"math"
)

//MaxArea ...
func MaxArea(height []int) int {
	//双指针对撞
	// 小的一边对求最大值没有影响，需要移动。
	// math.Min(x, y float64) float64
	max, left, right := 0, 0, len(height)-1
	for left < right {
		cur := (right - left) * int(math.Min(float64(height[left]), float64(height[right])))
		if cur > max {
			max = cur
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}
