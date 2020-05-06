/*
Package string ...
*/
package string

import (
	"strconv"
)

func addBinary(a string, b string) string {
	s := 0
	carry := 0
	res := ""
	la := len(a) - 1
	lb := len(b) - 1
	//When carry is 0 and la < 0 and lb < 0, stop the loop
	for la >= 0 || lb >= 0 || carry != 0 {
		//just continue to add number to s, check at the end of loop.
		s = carry
		if la >= 0 {
			s += int(a[la] - '0')
			la--
		}
		if lb >= 0 {
			s += int(b[lb] - '0')
			lb--
		}
		//Get new carry, and add remain to the result string.
		carry = s / 2
		res = string(s%2+'0') + res
	}
	return res
}

//cannot handle the situation that the number of bits is over 64...
func addBinary2(a string, b string) string {
	binA, _ := strconv.ParseInt(a, 2, 64)
	binB, _ := strconv.ParseInt(b, 2, 64)
	var answer int64
	var carry int64
	for binB != 0 {
		answer = binA ^ binB
		carry = (binA & binB) << 1
		binA, binB = answer, carry
	}
	return strconv.FormatInt(binA, 2)

}
