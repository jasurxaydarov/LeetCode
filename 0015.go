package main

import (
	"fmt"
)

func main() {
	str := "abcdf"
	k := 2
	res:=reverseStr(str, k)

	fmt.Println(res)
}

func reverseStr(s string, k int) string {
	start, end := "", ""
	res:=""
	for i := 0; i <k; i++ {
		start = start + string(s[i])
	}

	for i := k; i < len(s); i++ {
		end = end + string(s[i])
	}
	re := []rune(start)
	for i, j := 0, len(start)-1; i < j; i, j = i+1, j-1 {

		re[i], re[j] = re[j], re[i]
	}

	for _,w:=range re{

		res=res+string(w)
	}

	res=res+end


	return res
}
