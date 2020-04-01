/*
Package dp ...
*/
package dp

//SuperEggDrop K eggs and total N floors. Your goal is to know with certainty what the value of F is.
//What is the minimum number of moves that you need to know with certainty what F is, regardless of the initial value of F?
func SuperEggDrop(K int, N int) int {
	// K eggs, N floors
	// g(m, k) m moves and k eggs. g 代表最差情况下可以找到确定F 的层数
	// g(1, k) = 1 不论多少蛋移动一次确定 1 层
	// g(m, 1) = m 1个蛋最多确定 m 层
	// g(m, k) = 1 + g[m - 1][k - 1] + g[m - 1][k]
	//        当前移动 + 蛋碎了下面找 + 蛋没碎上面找
	// g(x-1, K) < N and g(x, K) >= N
	//横排代表的是move的次数，从左到右代表鸡蛋的个数，不管鸡蛋多少，一旦move满足条件就可以退出了
	g := make([][]int, N+1)
	for i := range g {
		g[i] = make([]int, K+1)
	}
	for m := 1; m < len(g); m++ {
		for k := 1; k < len(g[0]); k++ {
			g[m][k] = 1 + g[m-1][k-1] + g[m-1][k]
			if g[m][k] >= N {
				return m
			}
		}
	}
	return -1
}
