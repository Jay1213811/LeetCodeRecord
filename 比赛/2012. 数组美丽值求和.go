package main

import "fmt"

//如果i大于左边最大值,且小于右边最小值+2
//如果 i-1<i<i+1	+1
//否则+0

func sumOfBeauties(nums []int) int {
	n:=len(nums)
	LeftMax:=make([]int,n)
	RightMin:=make([]int,n)
	for i:=range LeftMax{
		LeftMax[i]=nums[0]
	}
	for j:=range RightMin{
		RightMin[j]=nums[len(nums)-1]
	}
	for i:=1;i<n;i++{
		LeftMax[i]=max2(LeftMax[i-1],nums[i-1])
	}
	for j:=n-2;j>=0;j--{
		RightMin[j]=min(RightMin[j+1],nums[j+1])
	}
	res:=0
	for i:=1;i<=n-2;i++{
		if nums[i]>LeftMax[i] && nums[i]<RightMin[i]{
			res+=2
		}else if nums[i]<nums[i+1] && nums[i]>nums[i-1]{
			res+=1
		}
	}
	fmt.Println(LeftMax)
	fmt.Println(RightMin)
	return  res

}
func max2(a,b int)int{
	if a>b{
		return a
	}else{
		return b
	}
}
func min(a,b int)int{
	if a<b{
		return a
	}else{
		return b
	}
}


func main()  {
	nums:=[]int{2,4,6,4}
	fmt.Println(sumOfBeauties(nums))
}