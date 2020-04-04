/*
Package array ...
*/
package array

//MoveZeroes : Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements
func MoveZeroes(nums []int) {

	nonzero := 0
	/* //O(N)
	   for i := 0; i < len(nums); i++ {
	       if nums[i] != 0 {
	           nums[nonzero++] = nums[i]
	           nonzero += 1
	       }
	   }
	   for i := nonzero; i < len(nums); i++ {
	       nums[i] = 0
	   }
	*/
	//Optimal operations swap times: number of nonzero)
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonzero], nums[i] = nums[i], nums[nonzero]
			nonzero++
		}
	}
}
