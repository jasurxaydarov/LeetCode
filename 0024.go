package main

import "fmt"

func main(){

	num:=[]int{1,1,2,3}

	res:=distributeCandies(num)
	fmt.Println(res)
}
func distributeCandies(candyType []int) int {
	res:=0
	
	check:=map[int]bool{}

	for _,w:=range candyType{


		check[w]=true
	}

	res=len(check)


	return res
}
