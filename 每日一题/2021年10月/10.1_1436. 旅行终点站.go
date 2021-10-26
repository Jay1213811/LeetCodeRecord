package main

import "fmt"

func destCity(paths [][]string) string {
	//用一个map来存储，每一站的起点终点。终点站的value为null
	travelMap:=make(map[string]string, len(paths))
	for i:=0;i<len(paths);i++{
		for j:=0;j<len(paths[0])-1;j++{
			begin:=paths[i][j]
			end:=paths[i][j+1]
			_,ok:=travelMap[begin]
			if !ok{
				travelMap[begin]=end
			}
		}
	}
	for _,v:=range(paths){
		for _,country:=range v{
			_,ok:=travelMap[country]
				if !ok{
					return country
				}
			}

	}
	return ""
}
func main()  {
	path:=[][]string{{"London","New York"},{"New York","Lima"},{"Lima","Sao Paulo"}}
	fmt.Print(destCity(path))
}