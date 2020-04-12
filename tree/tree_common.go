/*
Package tree ...
*/
package tree

//Node Definition for a binary tree node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

//Return Max(x, y)
func maxVal(x, y int) int {
	if x > y {
		return x
	}
	return y
}
