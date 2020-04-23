/*
Package array ...
*/
package array

func search(nums []int, target int) int {
	//binary answer
	if len(nums) == 0 {
		return -1
	}
	lo, hi := 0, len(nums)-1
	for lo+1 < hi {
		mid := lo + (hi-lo)/2
		//only two conditions: 1. the middle one is bigger than lo, 2, is smaller than lo.
		if nums[mid] > nums[lo] {
			//target is located at left part <= && <=
			if nums[lo] <= target && target <= nums[mid] {
				hi = mid
			} else {
				lo = mid
			}
		} else {
			//target is located at right part <= && <; !!!here is <
			if nums[mid] <= target && target < nums[lo] {
				lo = mid
			} else {
				hi = mid
			}
		}
	}

	if nums[lo] == target {
		return lo
	}
	if nums[hi] == target {
		return hi
	}
	return -1
}
