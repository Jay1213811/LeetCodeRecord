/*
1.首先把整数N转换为数组
如112——[1,1,2]=1*10的平方+1*10的1次方+2*10的0次方
然后开始dfs得到所有可能的情况
[[1,1,2],[1,2,1],[2,1,1]]
2.判断是否为2的幂函数。return n&(n-1) == 0
3.通过回溯得到N数组的所有不重复排列且第一位数字不为0
回溯的返回条件len(path)==len(ArrayN)
for遍历得到所有数组结果

如果是
*/
package main

import (
	"fmt",
	"sort"
)
func reorderedPowerOf2(n int) bool {
    arr:=[]int{}
    for n!=0{
        ans=append(arr,n%10)
        n=n/10
    }
	can:=false

    sort.Ints(arr)
    var dfs func(arr []int,used []bool,num int,depth int)
    dfs=func(arr []int,used []bool,num int,depth int){
     	if	can==true{
				return 
			}
    
		for i:=0;i<len(arr);i++{
			if used[i] || (depth == 0 && arr[i] == 0) || (i != 0 && arr[i] == arr[i-1] && !used[i-1]) {
				continue
			}
			used[i]=true
			dfs(arr,used,nums*10+arr[i],len+1)
		}
 
    for i:=startIndex;i<len(ans);i++{
        if i>startIndex&&ans[i]==ans[i-1]{
            continue
        }
        if ans[0]==0{
            continue
        }
        path=append(path, ans[i])
        dfs(i+1,path)
        path=path[:len(path)-1]
    }
    return false
 }
    dfs(0,[]int{})
    if len(res)!=0{
        return true
    }
    return false



}
func coverArrayToInt(nums []int)int{
    res:=0
    length:=len(nums)
    for i:=0;i<length;i++{
        res+=nums[i]
    }
    return res
}
func judgeIsPowOfTow(n int)bool{
        return n&(n-1) == 0
}
func powof(n int)int{
    if n==0{
        return 1
    }else{
        return 10*powof(n-1)
    }
}

func main(){
	fmt.Println(reorderedPowerOf2(1))
}