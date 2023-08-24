package main

import "fmt"

func main() {
	grade := 33

	if grade >= 80 && grade <= 100 {
		fmt.Println("Nilai anda A")
	}else if grade >= 65 && grade <= 79 {
		fmt.Println("Nilai anda B")
	}else if grade >= 50 && grade <= 64 {
		fmt.Println("Nilai anda C")
	}else if grade >= 35 && grade <= 49 {
		fmt.Println("Nilai anda D")
	}else if grade >= 0 && grade <= 34 {
		fmt.Println("Nilai anda E")
	}else{
		fmt.Println("Nilai Invalid")
	}
}