package main

import "fmt"

func ganjilGenap() {
	for a := 0; a <= 10; a++ {
		if a%2 == 0 {
			fmt.Print(a)
			fmt.Println(" adalah bilangan genap")
		}else{
			fmt.Print(a)
			fmt.Println(" adalah bilangan ganjil")
		}
	}
}