package main

import "fmt"

func main() {
	str := "ab_a"
	fmt.Println("aaaaaaaaa")

	res := isPalindrome(str)
	fmt.Println(res)

}

func isPalindrome(s string) bool {

	if len(s) == 1 {
		return true
	}
	var a, b string

	for i := 0; i < len(s); i++ {
		if s[i] > 64 && s[i] < 91 || s[i] > 96 && s[i] < 123 || s[i] > 47 && s[i] < 58 {
			if s[i] > 64 && s[i] < 91 {
				res := s[i] + 32
				a += string(res)
			} else {
				a += string(s[i])
			}
		}
	}
	fmt.Println(a)
	for v := len(s) - 1; v > -1; v-- {

		if s[v] > 64 && s[v] < 91 || s[v] > 96 && s[v] < 123 || s[v] > 47 && s[v] < 58 {

			if s[v] > 64 && s[v] < 91 {
				res := s[v] + 32
				b += string(res)
			} else {
				b += string(s[v])
			}
		}
	}
	fmt.Println(b)

	if a == b {
		return true

	} else {
		return false
	}
}
