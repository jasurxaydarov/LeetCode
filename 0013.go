package main

import (
	"fmt"
	"strings"
)

func main() {

	var l string = "maalL"

	res := DetectCapitalUse(l)

	fmt.Println(res)
}

func detectCapitalUse(word string) bool {
	check := true
	if len(word) == 1 {
		return check
	}
	if word[0] > 96 && word[1] > 96 && word[0] < 123 && word[1] < 123 {
		fmt.Println("AAA")
		// else if word[0] > 96 && word[1] > 96 && word[0] < 123 && word[1] < 123
		for i := 2; i < len(word); i++ {

			if word[i] < 96 || word[i] > 123 {
				check = false
				return check
			}
		}

	} else if word[0] > 64 && word[1] > 64 && word[0] < 91 && word[1] < 91 {
		fmt.Println("bb")
		// else if word[0] > 64 && word[1] > 64 && word[0] < 91 && word[1] < 91

		for i := 2; i < len(word); i++ {

			if word[i] < 64 || word[i] > 91 {
				check = false
				fmt.Println("bb")
				return check
			}
		}

	} else if word[0] > 64 && word[1] > 96 && word[0] < 91 && word[1] < 123 {

		fmt.Println("d")
		for i := 2; i < len(word); i++ {

			if word[i] < 96 || word[i] > 123 {
				check = false
				fmt.Println("dd")
				return check
			}
		}

	} else {
		fmt.Println("false")
		return false
	}

	return check
}

func DetectCapitalUse(word string) bool {
	return word == strings.ToUpper(word) || word == strings.ToLower(word) || ((word[0] >= 'A' && word[0] <= 'Z') && word[1:] == strings.ToLower(word[1:]))
}
