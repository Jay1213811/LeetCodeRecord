package main



func maxArea(height []int) int {
	max:=0
	start,end:=0,len(height)-1
	for start<end{
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		}else{
			high=height[end]
			end--
		}
		temp:=width*high
		if(temp>max){
			max=temp
		}
	}
	return max
}

