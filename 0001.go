package main

import "fmt"

func main(){
	num:=[]int{1,1,2}

	fmt.Println("aaaaaaaaa")

	res:=removeDuplicates(num)
	fmt.Println(res)
}

	func removeDuplicates(nums []int) int {

		if len(nums) == 0 {
			return 0
		}
		k := 1
	
		for i := 1; i < len(nums); i++ {
			if nums[i] != nums[i-1] {
				nums[k] = nums[i]
				k++
			}
		}
	
		return k
	}