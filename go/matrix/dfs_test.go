package matrix

import (
	"testing"
)

func TestDfs(t *testing.T) {
	tests := []struct {
		name string
		grid [][]uint8
		want int
	}{
		{
			name: "Grid 1",
			grid: [][]uint8{
				{0, 0, 0, 0},
				{1, 1, 0, 0},
				{0, 0, 0, 1},
				{0, 1, 0, 0},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if have := Dfs(tt.grid); have != tt.want {
				t.Errorf("wrong output\nhave: %v \nwant: %v", have, tt.want)
			}
		})
	}
}
