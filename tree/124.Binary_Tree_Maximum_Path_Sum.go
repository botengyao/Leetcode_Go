/*
Package tree ...
*/
package tree

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
 Given a non-empty binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections.
The path must contain at least one node and does not need to go through the root.
*/
var maxp int

//MaxPathSum ...
func MaxPathSum(root *Node) int {
	maxp = root.Val
	maxChain(root)
	return maxp
}

func maxChain(root *Node) int {
	if root == nil {
		return math.MinInt32
	}

	left := maxChain(root.Left)
	right := maxChain(root.Right)

	this := max(root.Val, root.Val+left, root.Val+right)
	maxp = max(maxp, this, left, right, root.Val+left+right)

	return this
}

func max(first int, nums ...int) int {
	m := first
	for _, v := range nums {
		if v > m {
			m = v
		}
	}
	return m
}
