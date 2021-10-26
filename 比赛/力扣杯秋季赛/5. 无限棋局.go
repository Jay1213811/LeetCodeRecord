package main

import "fmt"

func gobang(pieces [][]int) string {
	temp:=[]int{}
	white:=[][]int{}
	black:=[][]int{}
	for i:=0;i<len(pieces);i++{
		//0 表示黑棋，1 表示白棋
		if pieces[i][2]==1{
			temp=append(temp,pieces[i][0])
			temp=append(temp,pieces[i][1])
			white=append(white,temp)
			temp=[]int{}
		}else{
			temp=append(temp,pieces[i][0])
			temp=append(temp,pieces[i][1])
			black=append(white,temp)
			temp=[]int{}
		}
	}
	fmt.Println("BLACK",black)
	fmt.Println("white",white)

	return "NONE"
}
func main()  {
	piece:=[][]int{{1,2,1},{1,4,1},{1,5,1},{2,1,0},{2,3,0},{2,4,0},{3,2,1},{3,4,0},{4,2,1},{5,2,1}}
	fmt.Println(gobang(piece))
}