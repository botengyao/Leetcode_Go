/*
Package tree ...
*/
package tree

/**
 * Definition for a binary tree node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 * }
 */

//GenerateTrees ...
func GenerateTrees(n int) []*Node {
	memo := make(map[[2]int][]*Node)
	memo[[2]int{0, 0}] = []*Node{nil}
	if n == 0 {
		return nil
	}
	return generate(1, n, memo)
}

func generate(start, end int, memo map[[2]int][]*Node) []*Node {
	if start > end {
		return []*Node{nil}
	}

	if setTree, ok := memo[[2]int{start, end}]; ok {
		return setTree
	}

	var currRes []*Node
	for i := start; i <= end; i++ {
		lefttrees := generate(start, i-1, memo)
		righttrees := generate(i+1, end, memo)
		for _, l := range lefttrees {
			for _, r := range righttrees {
				currRes = append(currRes, &Node{Val: i, Left: l, Right: r})
			}
		}
	}
	memo[[2]int{start, end}] = currRes
	return currRes
}
