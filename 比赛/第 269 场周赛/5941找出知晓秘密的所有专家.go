package 第_269_场周赛

import (
	"fmt"
	"sort"
)

func findAllPeople(_ int, meetings [][]int, firstPerson int) (ans []int) {
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][2] < meetings[j][2] }) // 按照时间排序

	haveSecret := map[int]bool{0: true, firstPerson: true} // 一开始 0 和 firstPerson 都知道秘密
	for i, m := 0, len(meetings); i < m; {
		g := map[int][]int{}
		time := meetings[i][2]
		// 遍历时间相同的会议。注意这里的 i 和外层循环的 i 是同一个变量，所以整个循环部分的时间复杂度是线性的
		for ; i < m && meetings[i][2] == time; i++ {
			v, w := meetings[i][0], meetings[i][1]
			g[v] = append(g[v], w) // 建图
			g[w] = append(g[w], v)
		}

		vis := map[int]bool{} // 避免重复访问专家
		var dfs func(int)
		dfs = func(v int) {
			vis[v] = true
			haveSecret[v] = true
			for _, w := range g[v] {
				if !vis[w] {
					dfs(w)
				}
			}
		}
		for v := range g {
			if haveSecret[v] && !vis[v] { // 从在图上且知道秘密的专家出发，DFS 标记所有能到达的专家
				dfs(v)
			}
		}
	}
	for i := range haveSecret {
		ans = append(ans, i) // 注意可以按任何顺序返回答案
	}
	return
}

func main()  {
	n:=6
	meetings:=[][]int{{1,2,5},{2,3,8},{1,5,10}}
	fmt.Println(findAllPeople(n,meetings,1))
}