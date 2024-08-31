package main

import "fmt"



func main(){


	nums:=[]int{0,1,0,3,12}
	moveZeroes(nums)
	
	
}

func moveZeroes(nums []int)  {
    
	for i,_:=range nums{

		for j:=i+1;j<len(nums);j++{

			if nums[i]==0{
				
				nums[i]=nums[j]
				nums[j]=0

			}
		}
	}

	fmt.Println(nums)
}