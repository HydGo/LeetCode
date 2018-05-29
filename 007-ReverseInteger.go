package main

import "fmt"

func main(){
	//test
	n :=12345679
	fmt.Printf("%d",reverse(n))
}



func reverse(x int) int {
	const INT_MAX = int(^uint32(0) >> 1)
	sum  := 0
	flag :=0
	if x < 0{
		flag = 1
		x = 0 - x
	}
	for ;x>0;{
		sum= sum * 10 + x%10;
		x = x/10;
	}
	if flag == 1{
		sum = 0 -sum
	}
	if sum >INT_MAX || sum < ^INT_MAX{
		sum = 0
	}
	return sum
}