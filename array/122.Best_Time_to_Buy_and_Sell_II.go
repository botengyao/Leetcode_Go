/*
Package array ...
*/
package array

//MaxProfit ...
func MaxProfit(prices []int) int {
	//find 上升区间
	res := 0
	for i := 0; i < len(prices); i++ {
		if i != 0 && prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
