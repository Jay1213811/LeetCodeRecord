package main
/*
输入：version1 = "1.01", version2 = "1.001"
输出：0
解释：忽略前导零，"01" 和 "001" 都表示相同的整数 "1"
*/

/*
思路:双指针法
首先比较.前的数字大小
然后比较.后的数字大小
通过对小数点后的数字进行转换处理 y = y*10 + int(version2[j]-'0')
*/

func compareVersion(version1, version2 string) int {
	n, m := len(version1), len(version2)
	i, j := 0, 0
	//开始比较。其实就是分两次
	for i < n || j < m {

		ResultVersion1:=0
		for ; i < n&& version1[i] != '.'; i++ {
			ResultVersion1 =  ResultVersion1*10 + int(version1[i]-'0')
		}
		//如果是.号了跳出
		i++
		ResultVersion2:=0
		for ;j<m && version2[j]!='.';j++{
			ResultVersion2 =ResultVersion2*10+int(version2[j]-'0')
		}
		j++
		if ResultVersion1>ResultVersion2{
			return 1
		}else if ResultVersion1<ResultVersion2{
			return -1
		}
	}
	return 0


}


