package main

import "fmt"

func exist(board [][]byte, word string) bool {
	row := len(board)
	if row == 0 {
		return false
	}
	col := len(board[0])

	if col == 0 {
		return false
	}
	n := len(word)
	if n == 0 {
		return false
	}
	// r,c 当前搜索路径下标
	//idx 当前需要匹配的字符下标
	var dfs func(r, c, idx int) bool
	dfs = func(r, c, idx int) bool {
		// 目标字符串匹配完成
		if idx == n {
			return true
		}
		//当前搜索路径是否满足限制条件，不满足则“剪枝”，回溯到上一步。
		if r < 0 || r >= row || c < 0 || c >= col || board[r][c] != word[idx] {
			return false
		}
		//记录当前字符，以便之后恢复现场
		tmp := board[r][c]
		//当前路径字符字符设为 '0',暂时标记走过路径
		board[r][c] = '$'
		//继续上，下，左，右搜索
		if dfs(r+1, c, idx+1) ||
			dfs(r-1, c, idx+1) ||
			dfs(r, c+1, idx+1) ||
			dfs(r, c-1, idx+1) {
			return true
		}
		// 搜索回溯，撤销走过标记，恢复现场
		board[r][c] = tmp
		return false
	}

	//从矩阵每个位置出发搜索一遍
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			flag:=dfs(i, j, 0)
			if flag {
				return true
			}
		}
	}
	return false
}

func main(){
	board:=[][]byte{{'a','b','c','e'},{'s','f','c','s'},{'a','d','e','e'}}
	word:="sfde"
	fmt.Println(exist(board,word))
}
