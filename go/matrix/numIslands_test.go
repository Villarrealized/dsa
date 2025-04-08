package matrix

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]string
		want int
	}{
		{
			name: "Grid 1",
			grid: [][]string{
				{"0", "1", "1", "1", "0"},
				{"0", "1", "0", "1", "0"},
				{"1", "1", "0", "0", "0"},
				{"0", "0", "0", "0", "0"},
			},
			want: 1,
		},
		{
			name: "Grid 2",
			grid: [][]string{
				{"1", "1", "0", "0", "1"},
				{"1", "1", "0", "0", "1"},
				{"0", "0", "1", "0", "0"},
				{"0", "0", "0", "1", "1"},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if have := numIslands(tt.grid); have != tt.want {
				t.Errorf("wrong output\nhave: %v \nwant: %v", have, tt.want)
			}
		})
	}
}
