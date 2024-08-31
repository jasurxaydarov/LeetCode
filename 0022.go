package main

import "fmt"

func main() {
	num := []int{1}
	res := findMaxConsecutiveOnes(num)
	fmt.Println(res)
}

func findMaxConsecutiveOnes(nums []int) int {

	check := []int{}
	count := 0

	for i, w := range nums {

		if w == 1 {
			count++
		} else {

			check = append(check, count)
			count = 0
		}
		if i == len(nums)-1 {
			check = append(check, count)

		}

	}

	s := 0

	for _, w := range check {

		if s < w {

			s = w
		}

	}

	return s
}
