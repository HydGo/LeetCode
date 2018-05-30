package main

import "fmt"

func main(){
	prices :=[]int{7,1,5,3,6,}
	fmt.Printf("%d",maxProfit(prices))
}

func maxProfit(prices []int) int {
	var sum int = 0
	for i := 1 ;i<len(prices);i++{
		if prices[i-1]<prices[i]{
			sum += (prices[i]-prices[i-1])
		}
	}
	return sum
}
