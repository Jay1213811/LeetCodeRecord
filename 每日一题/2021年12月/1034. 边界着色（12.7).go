package main

import "fmt"

//没看懂题TAT
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func colorBorder(grid [][]int, row, col, color int) [][]int {
	m, n := len(grid), len(grid[0])
	type point struct{ x, y int }
	borders := []point{}
	originalColor := grid[row][col]
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	var dfs func(int, int)
	dfs = func(x, y int) {
		vis[x][y] = true
		isBorder := false
		for _, dir := range dirs {
			nx, ny := x+dir.x, y+dir.y
			if !(0 <= nx && nx < m && 0 <= ny && ny < n && grid[nx][ny] == originalColor) {
				isBorder = true
			} else if !vis[nx][ny] {
				vis[nx][ny] = true
				dfs(nx, ny)
			}
		}
		if isBorder {
			borders = append(borders, point{x, y})
		}
	}
	dfs(row, col)

	for _, p := range borders {
		grid[p.x][p.y] = color
	}
	return grid
}

func main(){
	grid:=[][]int{{1,1},{1,2}}
	row,col,color:=0,0,3
	fmt.Println(colorBorder(grid,row,col,color))
}