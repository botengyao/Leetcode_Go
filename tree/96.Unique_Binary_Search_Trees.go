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
/*
//Memorization
func numTrees(n int) int {
    memo := make(map[int]int)
    memo[0] = 1
    return findnum(n, memo)
}

func findnum(n int, memo map[int]int) int {
    if n == 1 {
        return 1
    }
    if num, ok := memo[n]; ok {
        return num
    }
    res := 0
    for i := 1; i <= n; i++ {
        res = res + findnum(i - 1, memo) * findnum(n - i, memo)
    }
    memo[n] = res
    return res
}
*/

//DP
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] = dp[i] + dp[j-1]*dp[i-j]
		}
	}
	return dp[len(dp)-1]
}
