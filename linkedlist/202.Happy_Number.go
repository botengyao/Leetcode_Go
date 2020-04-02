/*
Package linkedlist ...
*/
package linkedlist

//IsHappy 龟兔赛跑法，判断是否有环 O(logN) time O(1) space
func IsHappy(n int) bool {
	slow := n
	fast := getNext(n)
	for fast != slow && fast != 1 {
		fast = getNext(getNext(fast))
		slow = getNext(slow)
	}
	return fast == 1
}

func getNext(n int) int {
	var nnum int
	for n != 0 {
		remain := n % 10
		nnum += remain * remain
		n = n / 10
	}
	return nnum
}

/*
//IsHappy Hashset O(logN) space O(logN) time
func IsHappy(n int) bool {
    mem := make(map[int]struct{})
    var nnum int
    for {
        if _, ok := mem[n]; ok {
            return false
        } else {
            mem[n] = struct{}{}
            for n != 0 {
                remain := n % 10
                nnum += remain * remain
                n = n / 10
            }
            if nnum == 1 {
                return true
            }
            n = nnum
            nnum = 0
        }
    }
    return false
}*/
