/*
Package array ...
*/
package array

//LargestRectangleArea Given n non-negative integers representing the histogram's bar height where the width of each bar is 1,
//find the area of largest rectangle in the histogram.
func LargestRectangleArea(heights []int) int {
	//Input: [2,1,5,6,2,3]
	//Output: 10
	//思路：单调栈，存下标，从后向前计算
	// 2         (1 < 2)
	// pop(2) append(1)
	// 1 5 6
	// 1 5 6 2				append到2的时候发现 (2 < 6), 此时进入另一个分支，注意这是 i已经加1
	// 1 5 -     6 * 1		此时stack.pop()为 6， 6 * (i - stack[-1] - 1)
	// 1 5 - 2   (2 < 5)	继续判断发现2仍然 < 5
	// 1 - -     5 * 2		此时stack.pop()为 5， 5 * (i - stack[-1] - 1)
	// 1 - - 2				append 2
	// 1 - - 2 3			append 3
	// i > len(heights) 	继续处理stack中剩余变量
	// 1 - - 2 3 i			此时 i 在队尾
	// 1 - - 2 = i          stack.pop() = 3, 3 * (i - stack[-1](值是 4) - 1)
	// 1 = = = = i          stack.pop() = 2, 2 * (i - stack[-1](值是 1) - 1)
	// - - - - - i			stack.pop() = 1, 1 * (i) stack 空了, 乘以 i 就好了
	var stack []int
	stack = make([]int, 0)
	i, cur := 0, 0
	max := 0
	for i < len(heights) {
		if len(stack) == 0 || heights[i] >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
			i++
		} else {
			cur = heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				max = maxVal(max, cur*i) //可以初始化时stack = [-1] 这里就不需要判断了
			} else {
				max = maxVal(max, cur*(i-stack[len(stack)-1]-1))
			}
		}
	}
	for len(stack) > 0 {
		cur = heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			max = maxVal(max, cur*i)
		} else {
			max = maxVal(max, cur*(i-stack[len(stack)-1]-1))
		}
	}
	return max

}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
