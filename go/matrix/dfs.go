package matrix

type Matrix [][]uint8

func Dfs(grid Matrix, row int, col int, visited Matrix) int {
	ROWS := len(grid)
	COLS := len(grid[0])

	if min(row, col) < 0 ||
		row == ROWS ||
		col == COLS ||
		visited[row][col] == 1 ||
		grid[row][col] == 1 {
		return 0
	}

	if row == ROWS-1 && col == COLS-1 {
		return 1
	}

	visited[row][col] = 1

	count := 0
	count += Dfs(grid, row+1, col, visited)
	count += Dfs(grid, row-1, col, visited)
	count += Dfs(grid, row, col+1, visited)
	count += Dfs(grid, row, col-1, visited)

	visited[row][col] = 0
	return count
}
