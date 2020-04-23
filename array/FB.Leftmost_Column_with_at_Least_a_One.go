/*
Package array ...
*/
package array

/*
A binary matrix means that all elements are 0 or 1. For each individual row of the matrix, this row is sorted in non-decreasing order.

Given a row-sorted binary matrix binaryMatrix, return leftmost column index(0-indexed) with at least a 1 in it. If such index doesn't exist, return -1.
*/

/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 * type BinaryMatrix struct {
 *     Get(int, int) int
 *     Dimensions() []int
 * }
 */
//退出循环有两个条件，其中之一是row条数达到限制了，这时候需要check 是否是全0
func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	nm := binaryMatrix.Dimensions()
	index := -1
	for i, y := 0, nm[1]-1; i < nm[0] && y >= 0; {
		if binaryMatrix.Get(i, y) == 1 {
			index = y
			y--
		} else {
			i++
		}
	}
	return index
}
