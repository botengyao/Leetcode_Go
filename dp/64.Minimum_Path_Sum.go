/*
Package dp ...
*/
package dp

//MinPathSum ...
func MinPathSum(grid [][]int) int {
	//dp; it is fast if make the code be separated
	//the first colum
	for i := 1; i < len(grid); i++ {
		grid[i][0] += grid[i-1][0]
	}
	//the first row
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] += grid[0][j-1]
	}

	//other part
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] += minVal(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func minVal(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*
func minPathSum(grid [][]int) int {
    //dp
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if i == 0 && j > 0 {
                grid[i][j] += grid[i][j - 1]
            } else if j == 0 && i > 0 {
                grid[i][j] += grid[i - 1][j]
            } else if j == 0 && i == 0 {
                continue
            } else {
                grid[i][j] += minVal(grid[i - 1][j], grid[i][j - 1])
            }
        }
    }
    return grid[len(grid) - 1][len(grid[0]) - 1]
}

func minVal (x, y int) int {
    if x < y {
        return x
    }
    return y
}
*/
