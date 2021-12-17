package main

import "fmt"

func loudAndRich(richer [][]int, quiet []int) []int {
	grah := make([][]int, len(quiet))
	// 建图，没钱人指向有钱人
	for _, rich := range richer {
		grah[rich[1]] = append(grah[rich[1]], rich[0])
	}
	res := make([]int, len(quiet))
	// 初始化所有person的
	for i := range res {
		res[i] = -1
	}
	var dfs func(int)
	dfs = func(person int) {
		if res[person] != -1 {
			return
		}
		// 最初就是他自己
		res[person] = person
		// 遍历比他有钱的人
		for _, i := range grah[person] {
			// 比有钱人还有钱的人
			dfs(i)
			// 如果这个有钱人的安静值较小，则更新
			if quiet[res[i]] < quiet[res[person]] {
				res[person] = res[i]
			}
		}
	}
	for i := 0; i < len(quiet); i++ {
		dfs(i)
	}
	return res
}
func main(){
	richer:=[][]int{{1,0},{2,1},{3,1},{3,7},{4,3},{5,3},{6,3}}
	quiet:=[]int{3,2,5,4,6,1,7,0}
	fmt.Println(loudAndRich(richer,quiet))
}
