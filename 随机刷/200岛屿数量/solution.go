package leetcode
//Flood fill
func numIslands(grid [][]byte) int {
	res := 0
	for i:=0; i< len(grid[0]); i++ {
		for j:=0; j< len(grid); j++ {
			if grid[i][j] == 1 {
				fill(grid,i, j)
				res++
			}
		}
	}

	return res
}

func fill(g [][]byte, i, j int) {
	g[i][j] = 0
	if i+1 < len(g) {
		if g[i+1][j] == 1 {
			fill(g, i+1, j)
		}
	}
	if j+1 < len(g[i]) {
		if g[i][j+1] == 1 {
			fill(g, i, j+1)
		}
	}
	if i-1 >= 0 {
		if g[i-1][j] == 1 {
			fill(g, i-1, j)
		}
	}
	if j-1 >= 0 {
		if g[i][j-1] == 1 {
			fill(g, i, j-1)
		}
	}
}