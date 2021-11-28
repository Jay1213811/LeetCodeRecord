package 第_269_场周赛

func minimumDeletions(nums []int) int {
	if len(nums)==1{
		return 1
	}
	getIndex:=getMinAndMaxIndex(nums)
	//删除都从前面移除元素.那么返回的应该是靠后的元素下标值（下标值的最大值）+1
	//删除都从后面移除元素，那么返回的应该是靠前（下标值的最小值）的元素下标值+1
	//如果一个从前删除一个从后删除 前+1 后直接拿
	minIndex,maxIndex:=getIndex[0],getIndex[1]
	//删除都从前面移除元素
	a,b,c,d:=max(minIndex,maxIndex)+1,max(len(nums)-minIndex,len(nums)-maxIndex),minIndex+1+len(nums)-maxIndex,maxIndex+1+len(nums)-minIndex
	e:=min(a,b)
	f:=min(c,d)
	res:=min(e,f)
	return res
}
func getMinAndMaxIndex(nums []int)[]int {
	minIndex,maxIndex:=0,0
	res:=make([]int,2)
	min,max:=nums[0],nums[0]
	for i:=0;i<len(nums);i++{
		if nums[i]<=min{
			minIndex=i
			min=nums[i]
		}
		if nums[i]>=max{
			maxIndex=i
			max=nums[i]
		}
	}
	res[0]=minIndex
	res[1]=maxIndex
	return res
}
func min(a,b int)int  {
	if a<b{
		return  a
	}
	return  b
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
