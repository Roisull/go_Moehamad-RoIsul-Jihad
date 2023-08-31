package main

import "fmt"

func main() {
	fmt.Println(sum(12, 4))
	fmt.Println(sum(12, 4, 3))
	fmt.Println(sum(12, 4, 3, 6))
	fmt.Println(sum(12, 4, 5, 7))
}

func sum(data ...int) int{
	result := 0
	for _, v := range data{
		result += v
	}
	return result
}