package main

import "fmt"

func main(){
	var data = []int{1,3,2,5,10,8,7,9,4,6,12,15}
	var result []int

	for _, value := range data {
		if value%2 != 0{
			result = append(result, value)
		}
	}

	fmt.Print(result)
}