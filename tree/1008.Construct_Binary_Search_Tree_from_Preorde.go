/*
Package tree ...
*/
package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//BstFromPreorder ...
func BstFromPreorder(preorder []int) *Node {
	//[8,5,1,7,10,12]
	return preorderConstuct(0, len(preorder)-1, &preorder)
}

func preorderConstuct(start, end int, preorder *[]int) *Node {
	//Base case
	if end < start {
		return nil
	}

	root := &Node{(*preorder)[start], nil, nil}
	var next int
	//Stop at the first element that biger than the root
	for next = start; next <= end; next++ {
		if (*preorder)[next] > (*preorder)[start] {
			break
		}
	}
	root.Left = preorderConstuct(start+1, next-1, preorder)
	root.Right = preorderConstuct(next, end, preorder)
	return root
}
