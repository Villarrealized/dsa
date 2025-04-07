package main

import (
	"dsa/matrix"
	"fmt"
)

func dfs() {
	grid := [][]uint8{
		{0, 0, 0, 0},
		{1, 1, 0, 0},
		{0, 0, 0, 1},
		{0, 1, 0, 0},
	}
	visited := [][]uint8{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	res := matrix.Dfs(grid, 0, 0, visited)
	fmt.Println(res)
}

func main() {
	dfs()
}
