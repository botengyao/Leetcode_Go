/*
Package linkedlist ...
*/
package linkedlist

/*
You have a queue of integers, you need to retrieve the first unique integer in the queue.

Implement the FirstUnique class:

FirstUnique(int[] nums) Initializes the object with the numbers in the queue.
int showFirstUnique() returns the value of the first unique integer of the queue, and returns -1 if there is no such integer.
void add(int value) insert value to the queue.
*/

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type DoubleList struct {
	Size int
	Head *Node
	Tail *Node
}

type FirstUnique struct {
	HashNode  map[int]*Node
	HashCount map[int]int
	Double    *DoubleList
}

func (list *DoubleList) DeleteNode(node *Node) {
	prev := node.Prev
	next := node.Next
	if list.Head == node && list.Tail == node {
		list.Head = nil
		list.Tail = nil
	} else if list.Head == node && list.Tail != node {
		list.Head = next
		node.Prev = nil
	} else if list.Head != node && list.Tail == node {
		list.Tail = prev
		prev.Next = nil
	} else {
		prev.Next = next
		next.Prev = prev
	}
}

func (list *DoubleList) AddNode(node *Node) {
	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else {
		prev := list.Tail
		prev.Next = node
		node.Prev = prev
		list.Tail = node
	}
}

//NewFirstUnique 里面有map这种的话最好是加上 构造函数
func NewFirstUnique() FirstUnique {
	return FirstUnique{
		HashNode:  make(map[int]*Node),
		HashCount: make(map[int]int),
		Double: &DoubleList{
			Size: 0,
			Head: nil,
			Tail: nil,
		},
	}
}

func Constructor(nums []int) FirstUnique {
	f := NewFirstUnique()
	for _, num := range nums {
		f.Add(num)
	}
	return f
}

func (this *FirstUnique) ShowFirstUnique() int {
	if this.Double.Head == nil {
		return -1
	}
	return this.Double.Head.Val

}

func (this *FirstUnique) Add(value int) {
	if _, ok := this.HashCount[value]; ok {
		this.HashCount[value]++
		if node, ok := this.HashNode[value]; ok {
			this.Double.DeleteNode(node)
			delete(this.HashNode, value) //删除Map中的KEY
		}
	} else {
		this.HashCount[value] = 1
		node := &Node{value, nil, nil}
		this.Double.AddNode(node)
		this.HashNode[value] = node
	}
}

/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */

/*
   First Unique Number
Solution
You have a queue of integers, you need to retrieve the first unique integer in the queue.

Implement the FirstUnique class:

FirstUnique(int[] nums) Initializes the object with the numbers in the queue.
int showFirstUnique() returns the value of the first unique integer of the queue, and returns -1 if there is no such integer.
void add(int value) insert value to the queue.


Example 1:

Input:
["FirstUnique","showFirstUnique","add","showFirstUnique","add","showFirstUnique","add","showFirstUnique"]
[[[2,3,5]],[],[5],[],[2],[],[3],[]]
Output:
[null,2,null,2,null,3,null,-1]

Explanation:
FirstUnique firstUnique = new FirstUnique([2,3,5]);
firstUnique.showFirstUnique(); // return 2
firstUnique.add(5);            // the queue is now [2,3,5,5]
firstUnique.showFirstUnique(); // return 2
firstUnique.add(2);            // the queue is now [2,3,5,5,2]
firstUnique.showFirstUnique(); // return 3
firstUnique.add(3);            // the queue is now [2,3,5,5,2,3]
firstUnique.showFirstUnique(); // return -1
*/
