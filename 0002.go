package main

import "fmt"

func main(){

	nums:=[]int{4}

	res:=containsDuplicate(nums)
	fmt.Println(res)


}

func containsDuplicate(nums []int) bool {


	



	for i,v:=range nums{

		for j:=i+1;j<len(nums);j++{


			if v==nums[j]{

				return true
			}
		}

	}
    
	return false
}

