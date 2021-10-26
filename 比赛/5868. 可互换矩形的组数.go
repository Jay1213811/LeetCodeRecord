package main

import "fmt"

func interchangeableRectangles(rectangles [][]int) int64 {
	result:=make([]float64,0)
	for i:=0;i<len(rectangles);i++{
		for j:=0;j<len(rectangles[i])-1;j++{
			ans:=float64(rectangles[i][j+1])/float64(rectangles[i][j])

			result=append(result,ans)
			//fmt.Println("ans",ans)
		}
	}
	ans:=0
	countMap:=make(map[float64]int,0)
	for _,v:=range result{
		countMap[v]++
	}
	fmt.Println(countMap)
	for _,v := range countMap{
		if v>=2{
			ans+=count(v)
		}
	}

	return int64(ans)
}
func count(n int)  int{
	sum:=0
	for n>1{
		n=n-1
		sum+=n
	}

	return sum
}
func main()  {

	k:=[][]int{{4,8},{3,6},{10,20},{15,30}}
	j:=[][]int{{4,5},{7,8}}
	u:=[][]int{{4,2},{1,3},{4,1},{4,2},{2,4},{1,1},{1,1}}
	fmt.Println(interchangeableRectangles(u))
	fmt.Println(interchangeableRectangles(j))
	fmt.Println(interchangeableRectangles(k))
}