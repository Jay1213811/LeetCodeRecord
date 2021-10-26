package main

import (
	"fmt"
	"sort"
)

func numberOfWeakCharacters(properties [][]int) int {
	res:=0
	//排序
	n:=len(properties)
	sort.Slice(properties,func(i,j int)bool{
		if properties[i][0]==properties[j][0]{
			return  properties[i][1]<properties[j][1]
		}else{
			return properties[i][0]>properties[j][0]
		}

	})
	maxd:=properties[0][1]

	for i:=0;i<n;i++{
		if properties[i][1]<maxd{
			res++
			continue
		}else if properties[i][1]>maxd{
			maxd=properties[i][1]
		}
	}
	return res
}
func main()  {
	a:=[][]int{{1,5},{10,4},{4,3}}
	fmt.Println(numberOfWeakCharacters(a))
}