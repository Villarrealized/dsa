package matrix

/*
Given a 2D grid grid where '1' represents land and '0' represents water, count and return the number of islands.

An island is formed by connecting adjacent lands horizontally or vertically and is surrounded by water. You may assume water is surrounding the grid (i.e., all the edges are water).

Ex.
grid = [
    ["0","1","1","1","0"],
    ["0","1","0","1","0"],
    ["1","1","0","0","0"],
    ["0","0","0","0","0"]
  ]
*/

func numIslands(grid [][]string) int {
	rowCount := len(grid)
	columnCount := len(grid[0])
	islandCount := 0
	var dfs func(int, int)

	dfs = func(row int, col int) {
		if min(row, col) < 0 ||
			row == rowCount || col == columnCount ||
			grid[row][col] == "0" {
			return
		}

		grid[row][col] = "0"

		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row, col-1)
	}

	for r := range rowCount {
		for c := range columnCount {
			if grid[r][c] == "1" {
				dfs(r, c)
				islandCount += 1
			}
		}
	}

	return islandCount
}
