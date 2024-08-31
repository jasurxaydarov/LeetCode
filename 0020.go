package main

import "fmt"

func main() {

	j,a:="aA","aAAbbbb"

	res:=numJewelsInStones(j,a)

	fmt.Println(res)
}

func numJewelsInStones(jewels string, stones string) int {


	count:=0
	check := map[string]bool{}

	for _,w:=range jewels{
		check[string(w)]=true
	}

	for _,w:=range stones {

	if check[string(w)]{

		count++
	}

	

	}


	return count
}
