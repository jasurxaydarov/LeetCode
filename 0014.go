package main

import (
	"fmt"
	"strings"
)

func main(){

	s:="Let's take LeetCode contest"
	s=reverseWords(s)
	fmt.Println(s)
	
}

func reverseWords(s string) string {
	
	words:=strings.Split(s, " ")

	for i,word:=range words{

		words[i]=reversString(word)
	}

	res:=strings.Join(words," ")
	return res
}

func reversString(s string)string{

	var word =[]rune(s)

	for i,j:=0,len(s)-1;i<j;i,j=i+1,j-1{

		word[i],word[j]=word[j],word[i]
	}

	return string(word)
}