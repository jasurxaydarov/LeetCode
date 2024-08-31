
package main

import "fmt"

func main(){
	num:=[]int{3,2,2,3}

	fmt.Println("aaaaaaaaa")

	res:=removeElement(num,3)
	fmt.Println(res)
}

func removeElement(nums []int, val int) int {
	 
	res:=0

	
		for _,v:=range nums {
			if v != val{
			
				nums[res]=v
				res ++
			}
		}

		return res
}