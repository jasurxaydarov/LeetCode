package main

import (
	
	"fmt"
	
)


func main(){

	res:=addDigits(111)

	fmt.Println(res)

}

func addDigits(num int) int {
    for num>=10 {

	res:=0
	for num>0{

		res+=num%10
		num/=10
	}
		num=res
	}
	return num
}