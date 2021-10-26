package main

import "reflect"

func findRotation(mat [][]int, target [][]int) bool {
	//如果本身就相等
	if reflect.DeepEqual(mat,target){
		return true
	}
	for i:=0;i<4;i++{
		rotate(mat)
		if reflect.DeepEqual(mat,target){
			return true
		}
	}
	return false

}
//旋转90度
func rotate(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < l / 2; i++ {
		for j := i; j < l-i-1; j++ {
			matrix[i][j], matrix[j][l-1-i], matrix[l-1-i][l-1-j], matrix[l-1-j][i] = matrix[l-1-j][i], matrix[i][j], matrix[j][l-1-i], matrix[l-1-i][l-1-j]
		}
	}
}
