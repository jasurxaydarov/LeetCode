package main

import "fmt"

func main() {
	s := "()(){}[][}"
	res:=isValid(s)
	fmt.Println(res)

}

func isValid(s string) bool {

	if len(s)==1{
		return false
	}
	
	check := true

	c := 1
	for i := 0; i < len(s)-1; i++ {

		if s[i] == '{' && s[c] == ']' ||s[i] == '{'&& s[c] == ')' {
			fmt.Println("1111111",s[i],s[c])
			check = false
			return check
		}

		if s[i] == '(' && s[c] == ']' || s[i] == '(' && s[c] == '}' {
			fmt.Println("2222222222",s[i],s[c])
			check = false
			return check
		}

		if s[i] == '[' && s[c] == '}' ||s[i] == '[' &&  s[c] == ')' {
			fmt.Println("333333333",s[i],s[c])
			check = false
			return check
		}
		c++
	}
	return check
}
