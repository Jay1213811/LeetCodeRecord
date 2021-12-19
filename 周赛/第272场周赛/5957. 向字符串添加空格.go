package 第272场周赛
func addSpaces(s string, spaces []int) string {
	res:=make([]byte,0)
	index:=0
	for i,v :=range s{
		if index<len(spaces) && spaces[index]==i{
			res=append(res,' ')
			index++
		}
		res=append(res,byte(v))
	}

	return string(res)
}
func main()  {
	s := "LeetcodeHelpsMeLearn"
	space:=[]int{8,13,15}
	print(addSpaces(s,space))
}