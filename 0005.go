package main

import "fmt"

func main(){

	nums:=[]int{1,2,3,1,2,3}
		res:=containsNearbyDuplicate(nums,2)
	fmt.Println(res)
}

func containsNearbyDuplicate(nums []int, k int) bool {
 
	check:=map[int]int{}
	
	for _,v:=range nums{

		check[v]++
		
	}

	
		if check[k]==2{

			return false
		}
	

	return true
}