package main

import "fmt"

/*
用一个map存储每一个数出现次数
然后遍历target出现一次减少一次，如果为0，增加次数+1
*/

func minimumSwitchingTimes(source [][]int, target [][]int) int {
	ResMap:=map[int]int{}
	res:=0
	for i:=0;i<len(source);i++{
		for j:=0;j<len(source[0]);j++{
			_,ok:=ResMap[source[i][j]]
			if(ok){
				ResMap[source[i][j]]++
			}else{
				ResMap[source[i][j]]=1
			}
		}
	}
	//遍历target出现一次减少一次，如果为0，增加次数+1
	for i:=0;i<len(target);i++{
		for j:=0;j<len(target[0]);j++{
			_,ok:=ResMap[target[i][j]]
			if(ok)&& ResMap[target[i][j]]>0{

				ResMap[target[i][j]]--
			}else{
				res++
			}
		}
	}
	return res


}
func main()  {
	source:=[][]int{{1,3},{5,4}}
	target:=[][]int{{1,1},{6,5}}
	fmt.Println(minimumSwitchingTimes(source,target))
}