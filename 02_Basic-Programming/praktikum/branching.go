package main

import "fmt"

func main() {
	// if, else if, else
	hour := 15
	if hour < 12 {
		fmt.Println("Selamat Pagi")
	}else if hour < 18{
		fmt.Println("Selamat Sore")
	}else{
		fmt.Println("Selamat Malam")
	}

	// switch
	var today int = 2
	switch today{
	case 1: fmt.Println("Today is Monday")
	case 2: fmt.Println("Today is Tuesday")
	default: fmt.Println("Invalid Date")
	}
}