package main

import "fmt"

func main() {

	nums := []int{10,10,10,10}

	res := isPossibleToSplit(nums)

	fmt.Println(res)

}



func isPossibleToSplit(nums []int) bool {

	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++

		if freq[num] > 2 {
			return false
		}
	}
	
	distinctCount := len(freq)
	fmt.Println(freq)

	if distinctCount >= len(nums)/2 {
		return true
	}
	
	return false
}
