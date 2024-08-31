package main

import "fmt"

func main() {
	s := []byte{'h', 'e', 'l', 'l'}

	reverseString(s)

}

func reverseString(s []byte) {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {

		s[i], s[j] = s[j], s[i]
	}

	fmt.Println(string(s))
}