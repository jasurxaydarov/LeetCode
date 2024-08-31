package main

import "fmt"

func main() {

	res := checkRecord("LL")
	fmt.Println(res)
}

func checkRecord(s string) bool {

	res := true

	if len(s) == 1 {
		return true
	}

	for i := 0; i < len(s); i++ {

			if s[i] == 'L' && s[i+1] == 'L' && s[i+2] == 'L' {
				fmt.Println("aaaa")
				res = false
				return res

			} else if s[i] == 'A' && s[i+1] == 'A' {
				fmt.Println("bb")

				res = false
				return res
			}
		}
	

	return res
}
