package main

import "fmt"

func decodeCiphertext(encodedText string, rows int) string {
	encodedArray:=[][]byte{}
	length:=len(encodedText)
	temp:=[][]int{}
	res:=[]byte{}
	// (0,0) (0,1) (0,2) (0,3)
	// (1,0) (1,1) (1,2) (1,3)
	// (2,0) (2,1) (2,2) (2,3)
	//结果是(0,0)(1,1)(2,2)+(0,1)(1,2)(2,3)
	//[['c','h',' ',' '],[' ','i','e',' '],[' ',' ','p','r']]
	cows:=length/rows	//行数
	i:=0
	for i<len(encodedText){
		j:=i+cows
		temp:=[]byte{}
		a:=[]byte(encodedText[i:j])
		i+=cows
		temp=append(temp,a...)
		encodedArray=append(encodedArray,temp)
	}
	//读斜对角线
	for i:=0;i<1;i++{
		//(0,0) (0,1) (0,2) (0,3)满足条件的才会进行一条线的添加
		//那么满足条件的要求是什么呢？列数为n那么对第一行没个数添加(n-1,n-1)操作后，最后得到的结果不越界代表满足条件
		for j:=0;j<cows;j++{

				temp=append(temp,[]int{i,j})

		}
	}

	for i:=0;i<len(temp);i++{
		x:=temp[i][0]
		y:=temp[i][1]
		for x1,y1:=x,y;x1<rows && y1<cows;{
			res = append(res,encodedArray[x1][y1])
			x1+=1
			y1+=1
		}
	}
	return deleteTailBlank(string(res))
}
func deleteTailBlank(str string) string {
	spaceNum := 0
	for i := len(str)-1; i >= 0; i-- {  // 去除字符串尾部的所有空格
		if str[i] == ' ' {
			spaceNum++
		} else {
			break
		}
	}
	return str[:len(str)-spaceNum]
}



func main()  {
	encodedText := " b  ac"
	rows := 2
	//" b  ac"
	//2
	fmt.Println(decodeCiphertext(encodedText,rows))

}