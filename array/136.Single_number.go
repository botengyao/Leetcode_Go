/*
Package array ...
*/
package array

//SingleNumber Given a non-empty array of integers, every element appears twice except for one. Find that single one.
//Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
func SingleNumber(nums []int) int {
	//1. Set
	//2. Bit Mapulation
	//3. Hash_table
	//No set in Go, so a map can be used. map[T]struct{} is better than map[T]bool because struct{} doesn't occupy memory
	set := make(map[int]struct{})
	var res int
	for _, v := range nums {
		if _, ok := set[v]; ok {
			delete(set, v)
		} else {
			set[v] = struct{}{}
		}
	}
	for k := range set {
		res = k
	}
	return res
}

/*
	// Bit Mapulation
	// a ^ 0 = a
    // a ^ a = 0
    // a ^ b ^ b = a ^ 0 = a
    var res int
    for _, num := range nums {
        res ^= num
    }
	return res
*/

/*
    m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}

	return -1
}*/
