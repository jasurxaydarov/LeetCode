package main

import (
	"fmt"
)

func main() {

	var ansomNote = "loveleetcode"

	res := firstUniqChar(ansomNote)
	fmt.Println(res)

}

func firstUniqChar(s string) int {

	if len(s)==1{
		return 0
	}
	check := map[string]int{}

	for _, v := range s {

		check[string(v)]++

	}

	for i,v:=range s{

		if check[string(v)]==1{

			return i
		}

	}
	return -1
}
