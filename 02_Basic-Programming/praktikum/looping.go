package main

import "fmt"

func main() {
	sentence := "Hello"
	for i := 0; i < len(sentence); i++{
		fmt.Printf(string(sentence[i]) + "-")
	} 
	fmt.Println("---------->>>>>>>>>")

	for pos, char := range sentence{
		fmt.Printf("character %c start at byte position %d\n", char, pos)
	}

	// looping for
	// for i := 0; i < 5; i++ {
	// 	fmt.Print(i)
	// }

	// looping with continue dan break
	// for j := 0; j < 5; j++{
	// 	if j == 1{
	// 		fmt.Println(j)
	// 		continue
	// 	}
	// 	if j > 3{
	// 		break
	// 	}
	// 	fmt.Println(j)
	// }
}