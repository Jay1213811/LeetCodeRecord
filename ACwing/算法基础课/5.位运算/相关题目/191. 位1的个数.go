package main

func hammingWeight(num uint32) int {
	res:=0
	for num!=0{
		num-=(num & -num)
		res++
	}
	return res
}