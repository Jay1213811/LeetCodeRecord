package main

func maxDistance(colors []int) int {
	maxDistance:=0
	for i:=0;i<len(colors);i++{
		for j:=i+1;j<len(colors);j++{
			if colors[i]!=colors[j]{
				distance:=abs(j-i)
				maxDistance=max(distance,maxDistance)
			}
		}
	}
	return maxDistance

}
func abs(a int) int{
	if a<0{
		return -a
	}
	return a
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}