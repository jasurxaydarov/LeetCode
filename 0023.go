package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	num := []int{0}
	res:=findRelativeRanks(num)

	fmt.Println(res)

}

func findRelativeRanks(score []int) []string {
	mp := map[int]string{}
	check := map[int]int{}
	res := []string{}

	for i, w := range score {

		check[w] = i

	}
	sort.Sort(sort.Reverse(sort.IntSlice(score)))
	
	for i, w := range score {
		i++

		if i == 1 {
			fmt.Println(i)

			mp[check[w]] = "Gold Medal"
		} else if i == 2 {
			mp[check[w]] = "Silver Medal"
		} else if i == 3 {
			mp[check[w]] = "Bronze Medal"
		} else {
			s := strconv.Itoa(i)
			mp[check[w]] = s
		}

	}

	for i,_:=range score{

		res = append(res, mp[i])
	}
	fmt.Println(mp)
	return res
}
