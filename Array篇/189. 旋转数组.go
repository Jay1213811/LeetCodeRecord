package main
//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
/*
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
*/

/*
思路1:
我们可以使用额外的数组来将每个元素放至正确的位置。
用 n 表示数组的长度，我们遍历原数组，
将原数组下标为 i 的元素放至新数组下标为 (i+k) mod n 的位置，最后将新数组拷贝至原数组即可。
 */


func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}
/*
思路2数组翻转
*/
func rotate2(nums []int, k int) {
	//重点是k不是k，k是开始旋转的位置和数组长度有关
	k %= len(nums)
	reverse(nums)
	reverse(nums[k:])
	reverse(nums[:k])
}
func reverse(nums []int)  {
	left,right:=0,len(nums)-1
	for left<right{
		nums[left],nums[right]=nums[right],nums[left]
		left++
		right--
	}
}
