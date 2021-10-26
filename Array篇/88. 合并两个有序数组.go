package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	result := make([]int, 0)
	var i,j int =0,0
	for i<m && j<n{
		if nums1[i]<nums2[j] || nums1[i]==nums2[j] {
			result =append(result ,nums1[i])
			i++
		}else{
			result =append(result ,nums2[j])
			j++
		}
	}
	for i<m{
		result =append(result ,nums1[i])
		i++
	}
	for j<n{
		result =append(result ,nums2[j])
		j++
	}

}
func main(){
	a:=[]int{1,2,3,0,0,0}
	b:=[]int{2,5,6}
      merge(a,3,b,3)
}