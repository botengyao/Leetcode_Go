/*
Package tree ...
*/
package tree

/**
 * Definition for a binary tree node.
 * type Node struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//DiameterOfBinaryTree Longest path
func DiameterOfBinaryTree(root *Node) int {
	var res int
	findheight(root, &res)
	return res
}

func findheight(root *Node, res *int) int {
	if root == nil {
		return 0
	}

	left := findheight(root.Left, res)
	right := findheight(root.Right, res)
	*res = maxVal(left+right, *res)
	return maxVal(right, left) + 1
}
