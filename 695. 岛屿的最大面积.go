package main

import "math"

func maxAreaOfIsland(grid [][]int) int {
	var ans float64
	var dfs func(grid [][]int, x, y int) int
	dfs = func(grid [][]int, x, y int) int {
		if x < 0 || y < 0 || x == len(grid) || y == len(grid[0]) || grid[x][y] != 1 {
			return 0
		}
		grid[x][y] = 0
		ans := 1
		ans += dfs(grid, x-1, y)
		ans += dfs(grid, x, y+1)
		ans += dfs(grid, x+1, y)
		ans += dfs(grid, x, y-1)
		return ans
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			ans = math.Max(float64(ans), float64(dfs(grid, i, j)))

		}
	}
	return int(ans)
}

func main() {

}
