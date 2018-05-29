package main

import "fmt"

func main(){
	a :=[]int{1,1,2}
	fmt.Printf("%d", removeDuplicates(a))
}



func removeDuplicates(nums []int) int {
	var i int = 0
	for j := 1;j<len(nums);j++{
		if nums[i] != nums[j]{
			i++
			nums[i] = nums[j]
		}
	}
	return i+1

}
