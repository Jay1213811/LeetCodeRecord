package main

import "fmt"

func stoneGameVII(stones []int) int {
	Bob,Alice:=0,0
	end:=len(stones)
	//Alice每次拿首或者尾
	start:=0
//	EndSum,StartSum:=0,0
	//先Alice拿 后Bob拿
	AliceTaken,BobTaken:=false,false
	for start<end{
		if AliceTaken==false{
			//拿尾比拿头好
			a:=sum(stones[start:end-1])
			b:=sum(stones[start+1:end])
			if a > b{
				end=end-1
				Alice+=sum(stones[start:end])
				end--
			}else if a<=b{
				start++
				Alice+=b
				fmt.Println("Alice=",Alice)
				start++
			}
			BobTaken=false
			AliceTaken=true
		}
		if BobTaken==false{
			c:=sum(stones[start:end-1])
			d:=sum(stones[start+1:end])
			//拿尾比拿头好
			if c>d{
				end=end-1
				Bob+=c
				end--
			}else if c<=d{
				start++
				Bob+=d
				start++
			}
			BobTaken=true
			AliceTaken=false

		}


	}
	return Alice-Bob

}
func sum(nums []int)int  {
	sum:=0
	for _,v:=range nums{
		sum+=v
	}
	return sum
}
func main()  {
	stones:=[]int{5,3,1,4,2}
	fmt.Println(stoneGameVII(stones))
}