package main

import "fmt"

func main(){
	matrix := [][]int {{1,2,3},{4,5,6},{7,8,9}}
	rotate(matrix);
	fmt.Printf("%d",matrix)
}


func rotate(matrix [][]int)  {
	a := len(matrix)
	for i := 0;i<a;i++{
		for j:=i ;j<a;j++{
			matrix[i][j],matrix[j][i]=matrix[j][i],matrix[i][j]
		}
	}
	for i := 0;i<a;i++{
		for j:=0 ;j<a/2;j++{
			matrix[i][j],matrix[i][a-1-j]=matrix[i][a-1-j],matrix[i][j]
		}
	}

}
