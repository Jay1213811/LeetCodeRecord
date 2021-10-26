package main
func construct2DArray(original []int, m int, n int) [][]int {
	res:=[][]int{}
	if len(original)!=m*n{
		return nil
	}
	path:=make([]int,0)
	for i:=0;i<len(original);i++{	// 0 1 2 3
		path=append(path,original[i])
		if len(path)==n{
			res=append(res,path)
			path=[]int{}
		}
	}
	return res
}