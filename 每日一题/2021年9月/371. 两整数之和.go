package main
func getSum(a int, b int) int {
	return DecToBin(a+b)
}
// 原理：除2取模是最低位
func DecToBin(n int) int{
	result := 0

	if n == 0 {
		return 0
	}

	for ;n > 0;n /= 2 {
		lsb := n % 2
		result = result*10+lsb
	}

	return result
}