package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "    day"
	fmt.Println("aaaaaaaaa")

	res := lengthOfLastWord(str)
	fmt.Println(res)

}


func lengthOfLastWord(sen string) int{
	
   res:=strings.Fields(sen)
   fmt.Println(res)
	res1:=0

	if len(res[0])==1{
		return 1
	}
 for i:=len(res)-1;i>0;i--{

	res1=len(res[i])
	fmt.Println(res1,"aaaaaa")

	return res1
 }


	return res1
}