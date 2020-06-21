package permutation

import "sort"

/*
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/

/*
Your input
[4,1,2,6,5,3]
Output
[4,1,3,2,5,6]
Expected
[4,1,3,2,5,6]
*/

func nextPermutation(nums []int) {
	first := 0
	//From right to left find the first decreasing one
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			first = i - 1
			break
		}
	}

	//From frist to the end, find the smallest value that is bigger than first
	second := len(nums) - 1
	for i := first + 1; i < len(nums); i++ {
		if nums[i] > nums[first] {
			second = i
		} else {
			break
		}
	}

	//Swap the first and second elements
	nums[first], nums[second] = nums[second], nums[first]

	//Sort the elements after the first elements
	sort.Ints(nums[first+1:])
}
