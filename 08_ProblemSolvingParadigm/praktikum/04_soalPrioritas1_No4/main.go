package main

import "fmt"

// Mohon maaf pak saya masih belum terlalu paham dengan soal ini. 
// sehingga saya minta bantuan kepada chatGPT
// tapi saya akan berusaha memahaminya

// coba semua kemungkinan nilai x, y, z dari 1 hingga 100
func solveEquations(A, B, C int) (int, int, int) {
	for x := 1; x <= 100; x++ {
		for y := 1; y <= 100; y++ {
			for z := 1; z <= 100; z++ {
				if x+y+z == A && x*y*z == B && x*x+y*y+z*z == C {
					return x, y, z
				}
			}
		}
	}
	return -1, -1, -1
}

func main() {
	var A, B, C int
	fmt.Print("Masukkan nilai A, B, dan C (dipisahkan spasi): ")
	fmt.Scan(&A, &B, &C)

	x, y, z := solveEquations(A, B, C)
	if x == -1 {
		fmt.Println("No Solution")
	} else {
		fmt.Printf("%d %d %d\n", x, y, z)
	}
}
