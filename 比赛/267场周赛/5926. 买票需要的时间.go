package main

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int {
	  ans:=0
	  for tickets[k]>0{
		 // fmt.Println(tickets)
		  for i:=0;i<len(tickets);i++{
			  if tickets[i]>0{
				  tickets[i]--
				  ans++
				  if tickets[k]<=0{
					  return  ans
				  }
			  }else{
				  continue
			  }
		  }
	  }

	  return ans

}
func main()  {
	fmt.Println(timeRequiredToBuy([]int{84,49,5,24,70,77,87,8},3))
}
//[84,49,5,24,70,77,87,8]
//3
//预期154