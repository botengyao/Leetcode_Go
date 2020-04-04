/*
Package template implements data-driven templates for generating textual
output such as HTML.
....
*/
package main

import (
	array "Leetcode_Go/array"
	"fmt"
)

func main3() {
	nums := []int{2, 3, 5}
	fmt.Println(array.CombinationSum(nums, 8))  //[2 2 2 2] [2 3 3] [3 5]]
	fmt.Println(array.CombinationSum2(nums, 8)) //[[2 3 3] [3 5]]
}

func main() {
	notdoubleCap()
	doubleCap()
	var stack1 []int
	stack1 = append(stack1, 1)
	fmt.Println(stack1)
	stack2 := make([]int, 0)
	fmt.Println(stack2)
	fmt.Printf("Stack2 cap: %d; len: %d\n", cap(stack2), len(stack2)) //Stack2 cap: 0; len: 0

	stack2 = append(stack2, 1)
	fmt.Println(stack2)                                               //[1]
	fmt.Printf("Stack2 cap: %d; len: %d\n", cap(stack2), len(stack2)) //Stack2 cap: 1; len: 1
	stack2 = append(stack2, 2)
	fmt.Printf("Stack2 cap: %d; len: %d\n", cap(stack2), len(stack2)) //Stack2 cap: 2; len: 2
	stack2 = append(stack2, 3)
	fmt.Printf("Stack2 cap: %d; len: %d\n", cap(stack2), len(stack2)) //Stack2 cap: 4; len: 3
	stack2 = stack2[:len(stack2)-1]
	fmt.Printf("Stack2 cap: %d; len: %d\n", cap(stack2), len(stack2)) //Stack2 cap: 4; len: 2

}
func notdoubleCap() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)  //0, 1 //cap is 2
	y := append(x, 2) //0, 1, 2 // cap is doubled and underlying array is relocated
	fmt.Println(x)    //0, 1
	z := append(x, 3) //0, 1, 3
	fmt.Println(y, z) //[0 1 2] [0 1 3]
}

func doubleCap() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)  //cap is 4
	y := append(x, 3) //no location change to underlying array
	fmt.Println(y)    //[0 1 2 3]
	fmt.Println(x)    //[0 1 2]
	z := append(x, 4)
	fmt.Println(y, z) //[0 1 2 4] [0 1 2 4]
}

/*
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()
	a := []int{1, 2, 3}
	s1 := a
	s2 := a[:]
	fmt.Print(a)
	fmt.Print(s1)
	fmt.Print(s2)
	fmt.Println()
	fmt.Printf("%p\n", &s1)
	fmt.Printf("%p\n", &s2)
	fmt.Printf("%p\n", &a)

	s := []int{1, 2, 3}
	b := [3]int{4, 5, 6}
	ref := b[:]
	fmt.Println("Existing array:\t", ref)
	t := append(s, ref...)
	fmt.Println("Existing slice:\t", s)
	fmt.Println("New slice:\t", t)
	ref[1] = -1
	fmt.Println("New slice 2:\t", t)
	s = append(s, ref...)
	fmt.Println("Existing slice:\t", s)
	s = append(s, s...)

	fmt.Println("s+s:\t\t", s)
	fmt.Println("ss" == "ss")
*/
