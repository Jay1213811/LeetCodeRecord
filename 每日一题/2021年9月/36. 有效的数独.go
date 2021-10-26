/*
1.判断行
2.判断列
3.判断3*3小矩阵
第一个判断行[[1-3 x=0 m=3 y[0-6]n[3-9]]]
i:=0;i<3;i++	x=0 m=3 y=0 n=3
	j:=0;j<3;j++
board[i][j]
第一个判断列
第二个[x=0;m=3	y=3 n=6]
i:=0;i<3;i++
	j:=3;j<6;j++
第三个[x=0;m=3	y=6 n=9]
i:=0;i<3;i++
	j:=6;j<9;j++
第四个					[[4-6 x=3 m=6 y[0-6]n[3-9]]]
i:=3;i<6;i++[x=3;m=6	y=0 n=3]
	j:=0;j<3;j++
第五个[x=3;m=6	y=3 n=6]
i:=3;i<6;i++
	j:=3;j<6;j++
第六个[x=3;m=6	y=6 n=9]
i:=3;i<6;i++
	j:=6;j<9;j++
第七个[x=6;m=9	y=0 n=3]	[[7-9 x=6 m=9 y[0-6]n[3-9]]]
i:=6;i<9;i++
	j:=0;j<3;j++
第八个
i:=6;i<9;i++[x=6;m=9	y=3 n=6]
	j:=3;j<6;j++
第九个[x=6;m=9	y=6 n=9]
i:=6;i<9;i++
	j:=6;j<9;j++

[".",".",".",".","5",".",".","1","."],
[".","4",".","3",".",".",".",".","."],
[".",".",".",".",".","3",".",".","1"],
["8",".",".",".",".",".",".","2","."],
[".",".","2",".","7",".",".",".","."],
[".","1","5",".",".",".",".",".","."],
[".",".",".",".",".","2",".",".","."],
[".","2",".","9",".",".",".",".","."],
[".",".","4",".",".",".",".",".","."]


*/

package main
func isValidSudoku(board [][]byte) bool {
	m:=len(board)
	n:=len(board[0])
	flag1:=judgeLine(board,0,0,m,n)
	flag2:=judgeRow(board,0,0,m,n)
	flag3:=judgeThree(board)
	if flag1&&flag2&&flag3{
		return true
	}
	return false
}
//每遍历完一行或一列就初始化
func initTable(table map[byte]int) {
	for i := range table{
		table[i] = 0
	}
}
//判断行	控制二维数组大小便于去判断3*3
func judgeRow(board [][]byte,x,y,m,n int)bool{
	hashMap:=make(map[byte]int,0)
	//统计每一行
	for i:=x;i<m;i++{
		//initTable
		initTable(hashMap)
		for j:=y;j<n;j++{
			if board[i][j]>='0' && board[i][j]<='9'{
				num,ok:=hashMap[board[i][j]]	//看是否存在
				if (ok){
					num++
					return false	//不止出现一次
				}else{
					num=1
				}
			}
		}
	}
	return true
}
//判断列
func judgeLine(board [][]byte,x,y,m,n int)bool  {
	hashMap:=make(map[byte]int,0)
	//统计每一列
	for i:=x;i<m;i++{
		//initTable
		initTable(hashMap)
		for j:=y;j<n;j++{
			if board[j][i]>='0' && board[i][j]<='9' {
				num,ok:=hashMap[board[i][j]]	//看是否存在
				if (ok){
					num++
					return false	//不止出现一次
				}else{
					num=1
				}
			}
		}
	}
	return true
}
func judgeThree(board [][]byte)bool{
	//最上面3个3*3
	for y:=0;y<9;y+=3{
		for n:=3;n<=9;n+=3{
			if judgeRow(board,0,y,3,n)==false || judgeRow(board,3,y,6,n)==false||judgeRow(board,6,y,9,n){
				return false
			}
			if judgeLine(board,0,y,3,n)==false || judgeLine(board,3,y,6,n)==false||judgeLine(board,6,y,9,n){
				return false
			}
		}
	}
	return true
}