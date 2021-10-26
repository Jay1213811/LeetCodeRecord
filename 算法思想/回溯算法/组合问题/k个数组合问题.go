package main

import "fmt"

/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
*/

func combine(n int, k int) [][]int {
	res:=[][]int{}	//res存放符合条件结果的集合
	var dfs func(start int,path []int)
	dfs= func(start int, path []int) {
		if len(path)==k{
			//存放结果
			temp:=make([]int,len(path))
			copy(temp,path)
			res=append(res,temp)
			return
		}
		for i:=start;i<=n-(k-len(path))+1;i++{
			path=append(path,i)//处理节点
			dfs(i+1,path)	//backtracking(路径,选择列表)
			path=path[:len(path)-1]//回溯，撤销处理结果

		}
	}
	dfs(1,[]int{})
	//fmt.Println(res)
	return res

}
func main()  {
	fmt.Println(combine(4,2))
}
/*
递归参数:
result存放符合条件结果的集合
path存放符合条件的结果
n和k代表从【1，n】取有可能的k个组合
每次从集合中选取元素，可选择的范围随着选择的进行而收缩，调整可选择的范围，就是要靠
startIndex。
*/

