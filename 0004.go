package main

import "fmt"

func main(){
nums:=[]int{3,2,6,5,0,3}
res:=maxProfit(nums)
fmt.Println(res)
	
}
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {

            minPrice = price
		} else {
			profit := price - minPrice
    
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}