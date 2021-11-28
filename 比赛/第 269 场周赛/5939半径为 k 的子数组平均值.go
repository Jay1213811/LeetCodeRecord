package 第_269_场周赛

import "fmt"
//超时 要用前缀和

func getAverages(nums []int, k int) []int {
	if len(nums)<k{
		return []int{-1}
	}
	res:=make([]int,len(nums))
	prefixSum:=make([]int,len(nums)+1)
	prefixSum[0]=0
	for i:=0;i<len(nums);i++{
		prefixSum[i+1]=prefixSum[i]+nums[i]
	}

	//置-1
	for i:=0;i<k;i++{
		res[i]=-1
	}
	for i:=len(res)-1;i>=len(res)-k;i--{
		res[i]=-1
	}
	for i:=k;i<len(nums)-k;i++{
		sum:=prefixSum[i+k+1]-prefixSum[i-k]

		avg:=sum/(2*k+1)
		res[i]=avg
	}
	return res
}

func main()  {
	nums:=[]int{7,4,3,9,1,8,5,2,6}
	k:=3
	//nums:=[]int{100}
	//k:=4584768740

	fmt.Println(getAverages(nums,k))
}