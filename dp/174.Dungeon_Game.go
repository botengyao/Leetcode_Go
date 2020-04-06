/*
Package dp ...
*/
package dp

/*
The demons had captured the princess (P) and imprisoned her in the bottom-right corner of a dungeon. The dungeon consists of M x N rooms laid out in a 2D grid.
Our valiant knight (K) was initially positioned in the top-left room and must fight his way through the dungeon to rescue the princess.

The knight has an initial health point represented by a positive integer. If at any point his health point drops to 0 or below, he dies immediately.

Some of the rooms are guarded by demons, so the knight loses health (negative integers) upon entering these rooms;
other rooms are either empty (0's) or contain magic orbs that increase the knight's health (positive integers).

In order to reach the princess as quickly as possible, the knight decides to move only rightward or downward in each step.

Write a function to determine the knight's minimum initial health so that he is able to rescue the princess.

For example, given the dungeon below, the initial health of the knight must be at least 7 if he follows the optimal path RIGHT-> RIGHT -> DOWN -> DOWN.

-2 (K)	 -3	     3
-5	    -10	     1
10	     30	    -5 (P)
*/

//CalculateMinimumHP ...
func CalculateMinimumHP(d [][]int) int {
	//dp[i][j]代表进入此房间所需要的最少血量
	//倒着想 消耗完后必须有一滴血剩余
	//From bottom-right to up-left position
	m := len(d)
	n := len(d[0])
	down, right := 0, 0
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			//bottom-right position
			if j == n-1 && i == m-1 {
				//if it contains magic orbs, the minimum health that can move in is 1
				if d[i][j] >= 0 {
					d[i][j] = 1
					//after consuming the health to fight with demons, still 1 is needed.
				} else {
					d[i][j] = 1 - d[i][j]
				}
				//bottom row
			} else if i == m-1 && j < n-1 {
				//magic orbs is larger that the needed health to move out, just 1 is enough.
				if d[i][j] >= d[i][j+1] {
					d[i][j] = 1
					//not enough orbs or demons in the room (negtive), check the needed health to move out
				} else {
					d[i][j] = d[i][j+1] - d[i][j]
				}

				//right row
			} else if j == n-1 && i < m-1 {
				//magic orbs is larger that the needed health to move out, just 1 is enough.
				if d[i][j] >= d[i+1][j] {
					d[i][j] = 1
					//not enough orbs or demons in the room (negtive), Just use input = output - concusming
				} else {
					d[i][j] = d[i+1][j] - d[i][j]
				}
				//other parts. We need to check the right and down positon, and chose the minimum one.
			} else {
				//check if the orbs are enough or not.
				if d[i][j] >= d[i+1][j] {
					down = 1
					//not enough orbs or demons in the room (negtive). Just use input = output - concusming
				} else {
					down = d[i+1][j] - d[i][j]
				}

				if d[i][j] >= d[i][j+1] {
					right = 1
				} else {
					right = d[i][j+1] - d[i][j]
				}

				d[i][j] = minVal(down, right)
			}
		}
	}
	return d[0][0]
}
