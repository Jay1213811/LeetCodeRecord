package main

import "fmt"

/*
首先通过回溯决定第一个机器人的最佳路径
然后记录这条路的下标，最后把这些给清0
首先在0，0
每次两个选择
*/

func gridGame(grid [][]int) int64{
	res:=[][]int{}	//res存放符合条件结果的集合

	sum:=0
	var dfs func(x int,y int,grid [][]int,path []int)
	dfs= func(x int,y int,grid [][]int,path []int) {
		if x>len(grid)||y>len(grid[0]){
			temp:=[]int{}
			copy(temp,path)
			res=append(res,temp)
		}
		for x:=0;x<len(grid);x++{
			for y:=0;y<len(grid[0]);y++{
				sum=grid[x+1][y]+grid[x][y+1]
				path=append(path,)
				dfs(x+1,y+1,grid,path)
				path=path[:len(path)-1]
			}
		}
	}

}



func main()  {

	grid:=[][]int{{2,5,4},{1,5,1}}
	fmt.Println(gridGame(grid))
}