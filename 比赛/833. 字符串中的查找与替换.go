package main

import (
	"fmt"
	"sort"

	"strings"
)

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	res:=s
	sort.Ints(indices)
	reverse(indices)
	for i:=0;i<len(indices);i++{
		old:=s[indices[i]:indices[i]+len(sources[i])]
		new:=targets[i]
		if old==sources[i]{
			res=strings.Replace(res,old,new,1)
		}
	}
	return res
}
func reverse(nums []int)  []int{
	i,j:=0,len(nums)-1
	for i<j{
		nums[i],nums[j]=nums[j],nums[i]
		i++
		j--
	}
	return nums
}
func main()  {
	s:="ejvzndtzncrelhedwlwiqgdbdgctgubzczgtovufncicjlwsmfdcrqeaghuevyexqdhffikvecuazrelofjmyjjznnjdkimbklrh"
	sources:=[]string{"gc","tov","vy","re","ikv","lo","dw","iqgdbd","ue","kimbk","tgu","qd","ndt","elhe","crq","zn","ec","ff","bz","ej","ua","rh","lw","jj","mfd","cz","ufn","ex","cjl","vz","cr","agh","znnj"}
	targets :=[]string{"dxs","hn","vfc","wnr","tira","rko","oob","mlitiwj","zrj","onpp","ot","c","lm","hbsje","dgc","ruf","qi","h","vzn","ja","ow","te","km","imq","pia","ixp","xsod","m","eat","yf","lzu","dgy","dyrsc"}
	index:=[]int{25,35,60,77,69,79,15,19,58,92,27,64,4,11,51,7,72,67,30,0,74,98,17,85,48,32,38,62,43,2,9,55,87}
	fmt.Println(findReplaceString(s,index,sources,targets))
}

//"jayflmruflzuhbsjeoobkmmlitiwjdxsotvznixpghnxsodcieatwspiadgcedgyzrjvfcmchhtiraqiowzwnrrkofjmyimqdyrscdonpplte"




