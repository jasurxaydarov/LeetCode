package main

import (
	"fmt"
)

func main() {

	var ansomNote = "AB&BaaB"

	res := toLowerCase(ansomNote)
	fmt.Println(res)

}

func toLowerCase(s string) string {
    
	res:=""
	
	for _,w:=range s{

		if w>64&&w<91{

		res=res+string(w+32)

		}else{

			res=res+string(w)
		}
	}

	return res
}