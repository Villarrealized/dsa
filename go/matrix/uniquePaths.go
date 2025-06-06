package matrix

/*
You are given a binary matrix Grid where 0s represent land and 1s represent rocks that can not be traversed.

Return the number of unique paths from the top-left corner of Grid to the bottom-right corner such that all traversed cells are land cells. You may only move vertically or horizontally through land cells. For an individual unique path you cannot visit the same cell twice.

	Ex. grid = {
			{0, 0, 0, 0},
			{1, 1, 0, 0},
			{0, 0, 0, 1},
			{0, 1, 0, 0},
		}
*/
func uniquePaths(grid [][]uint8) int {
	rowLen := len(grid)
	colLen := len(grid[0])

	visited := make([][]uint8, rowLen)
	for r := range visited {
		visited[r] = make([]uint8, colLen)
	}
	return countPaths(grid, rowLen, colLen, 0, 0, visited)
}

func countPaths(grid [][]uint8, rowLen int, colLen int, row int, col int, visited [][]uint8) int {
	if min(row, col) < 0 ||
		row == rowLen ||
		col == colLen ||
		visited[row][col] == 1 ||
		grid[row][col] == 1 {
		return 0
	}

	if row == rowLen-1 && col == colLen-1 {
		return 1
	}

	visited[row][col] = 1

	count := 0
	count += countPaths(grid, rowLen, colLen, row+1, col, visited)
	count += countPaths(grid, rowLen, colLen, row-1, col, visited)
	count += countPaths(grid, rowLen, colLen, row, col+1, visited)
	count += countPaths(grid, rowLen, colLen, row, col-1, visited)

	visited[row][col] = 0
	return count
}
