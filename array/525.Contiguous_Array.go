/*
Package array ...
*/
package array

//FindMaxLength ...
func FindMaxLength(nums []int) int {
	//Hash_map
	//0 -1, 1 +1
	//[-1, 1, 1, -1, 1]
	//相等的话是两个位置的差值为0！！！
	memo := make(map[int]int)
	memo[0] = -1
	sum, res := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			sum--
		} else {
			sum++
		}
		if id, ok := memo[sum]; ok {
			res = maxVal(res, i-id)
		} else {
			memo[sum] = i
		}
	}
	return res
}
