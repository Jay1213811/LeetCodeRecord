package main

import "fmt"

//二维前缀和板子
func pre_sum(g [][]int)[][]int{
	sum:=make([][]int,len(g))
	for i:=range sum{
		sum[i]=make([]int,len(g[0]))
	}
	//边界情况【如果下标从1开始就不需要处理边界情况即第一行和第一列全0的情况】
	sum[0][0]=g[0][0]
	for i:=1;i<len(sum);i++{
		sum[i][0]=sum[i-1][0]+g[i][0] //第一列
	}
	for j:=1;j<len(sum[0]);j++{
		sum[0][j]=sum[0][j-1]+g[0][j]//第一行
	}
	//通式
	for i:=1;i<len(sum);i++{
		for j:=1;j<len(sum[0]);j++{
			//下面的小块=整个矩形-左边一大块-上面一大块+左上角多减了一次的块
			//g[i][j]=sum[i][j]-sum[i-1][j]-sum[i][j-1]+sum[i-1][j-1]
			sum[i][j]=g[i][j]+sum[i][j-1]+sum[i-1][j]-sum[i-1][j-1]
		}
	}
	return sum
}
func main(){
	g:=[][]int{{1,5,6,8},{9,6,7,3},{5,3,2,4}}
	sum:=pre_sum(g)//求矩阵和
	//求小矩阵之和 (x1,y1)-(x2,y2)围起来的矩阵和
	var get_sum func(x1,y1,x2,y2 int)int
	get_sum= func(x1, y1, x2, y2 int) int{
		if x1==0 && y1==0{
			return sum[x2][y2]
		}
		if x1==0{
			return sum[x2][y2]-sum[x2][y1-1]
		}
		if y1==0{
			return sum[x2][y2]-sum[x1-1][y2]
		}
		return sum[x2][y2]-sum[x2][y1-1]-sum[x1-1][y2]+sum[x1-1][y1-1]
	}
	res1:=get_sum(1,1,2,2)
	res2:=get_sum(0,1,1,3)
	fmt.Println(res1,res2) //18 35

}