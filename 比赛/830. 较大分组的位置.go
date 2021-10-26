package main

import "fmt"

func largeGroupPositions(s string) [][]int {
	res:=[][]int{}
	path:=[]int{}
	slow,fast:=0,1
	for fast<len(s){
		if s[slow]!=s[fast]{

			slow=fast;
			fast++

		}else{
			if fast-slow>=2 && (fast==len(s)-1 || s[fast]!=s[fast+1]){
				path=append(path,slow)
				path=append(path,fast)
				res=append(res,path)
				path=[]int{}
			}

			fast++

		}


	}
	return res
}
func main(){
	fmt.Println(largeGroupPositions("aaa"))
}