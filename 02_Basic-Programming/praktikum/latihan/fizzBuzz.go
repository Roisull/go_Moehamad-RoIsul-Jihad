package main

import "fmt"

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz")
		}else if i%5 == 0{
			fmt.Print("Buzz")
		}else{
			fmt.Print(i)
		}
		fmt.Print(" ")
	}

}