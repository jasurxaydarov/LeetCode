package main

import "fmt"

func main() {

	nums := []int{1, 2, 2, 3, 1, 4}
	res := maxFrequencyElements(nums)
	fmt.Println(res)

}

func maxFrequencyElements(nums []int) int {

	check := map[int]int{}
	max := 0

	res := 0
	for _, w := range nums {

		check[w]++

		if check[w] > max {
			max = check[w]
		}
	}

	for _, w := range check {

		if max == w {
			res += max
		}
	}

	return res
}
