/*
Package tree ...
*/
package tree

import (
	"container/list"
)

//BFS
func rightSideView(root *Node) []int {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	res := make([]int, 0)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			e := queue.Front()
			if i == length-1 {
				res = append(res, e.Value.(*Node).Val)
			}
			queue.Remove(e)

			if e.Value.(*Node).Left != nil {
				queue.PushBack(e.Value.(*Node).Left)
			}
			if e.Value.(*Node).Right != nil {
				queue.PushBack(e.Value.(*Node).Right)
			}
		}
	}
	return res
}
