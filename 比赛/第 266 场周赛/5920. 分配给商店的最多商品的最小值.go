package main

import "fmt"

func minimizedMaximum(n int, quantities []int) int {
	l,r:=1,max(quantities)
	var check func(x int)bool
	check=func(x int)bool{
		ans:=0
		for _,v:=range quantities{
			if v%x==0{
				ans+=v/x
			}else{
				ans+=v/x+1
			}
		}
		return ans>n
	}
	for l<r{
		mid := (l+r)>>1
		if check(mid){
			l=mid+1
		}else{
			r=mid
		}
	}
	return l
}
func max(nums []int)int{
	res:=nums[0]
	for _,v:=range nums{
		if v>res{
			res=v
		}
	}
	return res
}
func main()  {
	fmt.Println(minimizedMaximum(2,[]int{5,7}))
}