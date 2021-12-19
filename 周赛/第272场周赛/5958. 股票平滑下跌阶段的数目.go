package 第272场周赛
func getDescentPeriods(prices []int) int64 {
	n:=len(prices)
	res:=n
	slow,fast:=0,1
	for fast<n{
		if  prices[fast-1]-1==prices[fast]{
			res+=(fast-slow)
			fast++
		}else{
			slow=fast
			fast++
		}
	}
	return int64(res)
}

