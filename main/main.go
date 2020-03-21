package main

import (
	array "Leetcode_Go/array"
	"fmt"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	fmt.Println(array.ThreeSum(nums))
	array.ThreeSumClosest(nums, 1)
}
