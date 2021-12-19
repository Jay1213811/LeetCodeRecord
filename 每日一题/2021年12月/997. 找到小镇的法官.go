package main
//12.9
//计算每个点的出度和入度
//找到入度=n-1 出度=0的点
func findJudge(n int, trust [][]int) int {
	in:=make([]int,n)
	out:=make([]int,n)
	for i:=0;i<len(trust);i++{
		in[trust[i][1]-1]+=1
		out[trust[i][0]-1]+=1
	}
	for i:=0;i<n;i++{
		if in[i]==n-1 && out[i]==0{
			return i+1
		}
	}
	return -1
}