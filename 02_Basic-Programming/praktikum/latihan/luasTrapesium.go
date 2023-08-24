package main

import "fmt"

const p float64 = 0.5

func main() {
	a := 5.0
	b := 6.0
	h := 10.0
	var L float64 = p * (a + b) * h
	fmt.Println(L)
}