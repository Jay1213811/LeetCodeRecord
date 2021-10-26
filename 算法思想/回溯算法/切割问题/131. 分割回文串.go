package 切割问题
func partition(s string) [][]string {
	res:=[][]string{}   // 存储结果
	var backtracking func(startIndex int,path []string)
	backtracking= func(startIndex int,path []string){
		//  终止条件startIndex>=len(s)  如果小于说明里面有不是回文串不保存
		if startIndex>=len(s){
			//存结果
			temp:=make([]string,len(path))
			copy(temp,path)
			res=append(res,temp)
		}
		//循环遍历
		for i:=startIndex;i<len(s);i++{
			//处理（首先通过startIndex和i判断切割的区间，进而判断该区间的字符串是否为回文，若为回文，则加入到path，否则继续后移，找到回文区间）（这里为一层处理）
			if judgeIsPartition(s,startIndex,i){
				path=append(path,s[startIndex:i+1])
			}else{
				continue
			}
			//递归
			backtracking(i+1,path)
			//回溯
			path=path[:len(path)-1]
		}
	}

	backtracking(0,[] string{})
	return res

}
//判断回文串
func judgeIsPartition(s string,startIndex,endIndex int)bool{
	start,end:=startIndex,endIndex
	for start<end{
		if s[start]!=s[end]{
			return false
		}else{
			start++
			end--
		}
	}
	return true
}