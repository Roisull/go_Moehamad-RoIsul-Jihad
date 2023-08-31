package main

import "fmt"

func operatorEksperession() {
	// operand dan operator
	x := 1 + 2
	fmt.Println(x)

	// expression
	a := 5
	b := 6
	c := a + b
	fmt.Println(c)

	// LUAS SEGITIGA
	alas := 10
	tinggi := 15
	luas := (alas * tinggi) / 2
	fmt.Println(luas)

	// string operation
	helloworld := "hello" + " " + "world"
	fmt.Println(helloworld)
}