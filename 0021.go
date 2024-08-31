package main

import (
	"fmt"
	"strings"
)

func main() {

	a := "Bob hit a ball, the hit, BALL flew far after it was hit."
	b := []string{"hit"}
	res := mostCommonWord(a, b)
	fmt.Println(res)
}

func mostCommonWord(paragraph string, banned []string) string {
	check := map[string]int{}
	banedSet := map[string]bool{}
	re := ""
	for i := 0; i < len(paragraph); i++ {

		if paragraph[i] < 64 && paragraph[i] < 91 && paragraph[i] < 97 && paragraph[i] < 123 && paragraph[i]==' '{

			re = re + string(" ")
			continue
		}
		re = re + string(paragraph[i])
	}

	paragraph = strings.ToLower(re)

	res := strings.Split(paragraph, " ")

	for _, w := range banned {

		banedSet[w] = true

	}
	
	for _, w := range res {

		if !banedSet[w] {
			if w != " " {
				check[w]++
			}
		}
	}

	fmt.Println(check)
	maxCount := 0
	var mostCommon string
	for word, count := range check {
		if count > maxCount {
			maxCount = count
			mostCommon = word
		}
	}

	return mostCommon

	

}
