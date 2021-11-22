package main

import (
	"fmt"
	"strconv"
)

func isRectangleCover(rectangles [][]int) bool {
	//初始化定义大矩形的四个顶点
	left:=rectangles[0][0]
	bottom:=rectangles[0][1]
	right:=rectangles[0][2]
	top:=rectangles[0][3]
	areaOfMaxRectangles:=0
	areaOfAllRectangles:=0
	hashMap:=make(map[string]int,0)
	//1.满足大矩形四个顶点出现1次.且其它所有点出现2次或者4次.45
	//2.n个小矩形面积之和等于大矩形
	for i:=0;i<len(rectangles);i++{
		A,B,C,D:=rectangles[i][0],rectangles[i][1],rectangles[i][2],rectangles[i][3]
		left=min(left,A)
		bottom=min(bottom,B)
		right=max(right,C)
		top=max(top,D)
		APoint:=strconv.Itoa(A)+","+strconv.Itoa(B)
		BPoint:=strconv.Itoa(A)+","+strconv.Itoa(D)
		CPoint:=strconv.Itoa(C)+","+strconv.Itoa(B)
		DPoint:=strconv.Itoa(C)+","+strconv.Itoa(D)
		hashMap=countPoint(hashMap,APoint)
		hashMap=countPoint(hashMap,BPoint)
		hashMap=countPoint(hashMap,CPoint)
		hashMap=countPoint(hashMap,DPoint)
		areaOfAllRectangles+=caculateArea(A,B,C,D)
	}
	LeftPoint:=strconv.Itoa(left)+","+strconv.Itoa(bottom)
	BottomPoint:=strconv.Itoa(left)+","+strconv.Itoa(right)
	RightPoint:=strconv.Itoa(right)+","+strconv.Itoa(bottom)
	TopPoint:=strconv.Itoa(top)+","+strconv.Itoa(right)
	if hashMap[LeftPoint]!=1 || hashMap[BottomPoint]!=1 ||hashMap[RightPoint]!=1||hashMap[TopPoint]!=1{
		return false
	}
	for i,v:=range hashMap{
		println(i,v)
		if v!=2 && v!=4 && v!=1{
			return false
		}
	}
	//条件2
	areaOfMaxRectangles=caculateArea(left,bottom,right,top)
	if areaOfMaxRectangles!=areaOfAllRectangles{
		return false
	}
	return true
}
func countPoint(m map[string]int,str string)map[string]int{
	_,ok:=m[str]
	if ok{
		m[str]+=1
	}else{
		m[str]=1
	}
	return m
}
func caculateArea(left,bottom,right,top int)int{
	return (top-bottom)*(right-left)
}
func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}
func main()  {

	rectangles:=[][]int{{0,0,1,1},{0,0,2,1},{1,0,2,1},{0,2,2,3}}
	fmt.Println(isRectangleCover(rectangles))
}
