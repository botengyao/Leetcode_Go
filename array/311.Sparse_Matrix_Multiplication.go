/*
Package array ...
*/
package array

func multiply(A [][]int, B [][]int) [][]int {
	la := len(A)
	lb := len(B[0])
	res := make([][]int, la)
	for i := 0; i < la; i++ {
		row := make([]int, lb)
		res[i] = row
	}
	for i := 0; i < la; i++ {
		for j := 0; j < lb; j++ {
			temp := 0
			for k := 0; k < len(B); k++ {
				temp = temp + A[i][k]*B[k][j]
			}
			res[i][j] = temp
		}
	}
	return res
}
