/*
Package array ...
*/
package array

/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

//Int32Max Max int 0111 1111 1111 11..
const Int32Max int = int(^uint(0) >> 1)

//Int32Min Min int 1000 0000 0000 00..
const Int32Min int = ^Int32Max

//MaxSubArray Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
func MaxSubArray(nums []int) int {
	//sliding window X
	//dp 以当前数结束最大的结果
	//divide and conquer
	//const UINT32_MIN uint32 = 0
	//const UINT32_MAX uint32 = ^uint32(0)
	//const Int32Max int = int(^uint(0) >> 1)
	//const Int32Min int = ^INT32_MAX
	max := nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+max > nums[i] {
			max = nums[i] + max
		} else {
			max = nums[i]
		}
		if max > res {
			res = max
		}
	}
	return res
}

//MaxSubArrayDC divide and Conquer
func MaxSubArrayDC(nums []int) int {
	return divide(0, len(nums)-1, nums)
}

func divide(start, end int, nums []int) int {
	if end <= start {
		return nums[end]
	}

	mid := start + (end-start)/2
	maxLeft := divide(start, mid, nums)
	maxRight := divide(mid+1, end, nums)
	maxMid := divideCross(start, end, mid, nums)

	if maxLeft >= maxRight && maxLeft >= maxMid {
		return maxLeft
	} else if maxRight >= maxLeft && maxRight >= maxMid {
		return maxRight
	} else {
		return maxMid
	}
}

func divideCross(start, end, mid int, nums []int) int {
	cur := 0
	rightM := Int32Min
	leftM := Int32Min
	for i := mid + 1; i <= end; i++ {
		cur += nums[i]
		if cur > rightM {
			rightM = cur
		}
	}
	cur = 0
	for j := mid; j >= 0; j-- {
		cur += nums[j]
		if cur > leftM {
			leftM = cur
		}
	}
	return leftM + rightM
}
