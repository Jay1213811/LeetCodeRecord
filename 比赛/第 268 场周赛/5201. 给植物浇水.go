package main

import "fmt"

func wateringPlants(plants []int, capacity int) int {
	maxWater:=capacity
	need:=0
	//模拟
	for i:=0;i<len(plants);{
		capacity=capacity-plants[i]
			if capacity>=0{
				//如果水够直接加
				need++
				i++
			}else{
				//水不够了
				need=need+i
				capacity=maxWater
				need=need+i+1
				capacity=capacity-plants[i]
				i++
			}
		}

	return  need

}


func main(){
	plants:=[]int{2,2,3,3}
	capacity:=5
	fmt.Println(wateringPlants(plants,capacity))
}