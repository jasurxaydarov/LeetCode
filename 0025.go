package main

import "fmt"


func main(){

num1:=[]int{1}
num2:=[]int{1}

	res:=addedInteger(num1,num2)

	fmt.Println(res)

}

func addedInteger(nums1 []int, nums2 []int) int {
    
	max1,max2:=0,0

	res:=0
	if nums1[0]>nums2[0]{
		for _,w:=range nums1{

			if max1<w{
				max1=w
			}

		}
		for _,w:=range nums2{

			if max2<w{
				max2=w
			}

		}

		fmt.Println(max1)
		fmt.Println(max2)

		res=max2-max1
	}else if  nums1[0]<nums2[0]{

		for _,w:=range nums1{

			if max1<w{
				max1=w
			}

		}
		for _,w:=range nums2{

			if max2<w{
				max2=w
			}

		}

		fmt.Println(max1)
		fmt.Println(max2)

		res=max2-max1

	}else if  nums1[0]==nums2[0]{

		return 0
	}

	

	return res
}

