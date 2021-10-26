package main


func combinationSum(candidates []int, target int) [][]int {
	res:=[][]int{}
	path:=[]int{}
	var backtracking func(candidates []int,index int,startIndex int)
	backtracking=func(candidates []int,index int,startIndex int){
		if index>=target{
			if index==target{
				temp:=make([]int,len(path))
				copy(temp,path)
				res=append(res, temp)
				return
			}
			return
		}
		for i:=startIndex;i<len(candidates);i++{
			index+=candidates[i]
			path=append(path, candidates[i])
			backtracking(candidates,index,i)
			index-=candidates[i]//回溯
			path=path[:len(path)-1]
		}
	}
	backtracking(candidates,0,0)
	return res

}