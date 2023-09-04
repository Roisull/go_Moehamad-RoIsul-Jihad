package main

import "fmt"

// Mohon maaf pak saya masih belum terlalu paham dengan soal ini. 
// sehingga saya minta bantuan kepada chatGPT
// tapi saya akan berusaha memahaminya

// coba semua kemungkinan nilai x, y, z dari 1 hingga 100
func SimpleEquations(a, b, c int) {
	for x := 1; x <= a; x++ {
        for y := 1; y <= a; y++ {
            for z := 1; z <= a; z++ {
                if x+y+z == a && x*y*z == b && x*x+y*y+z*z == c {
                    fmt.Printf("%d %d %d\n", x, y, z)
                    return
                }
            }
        }
    }
    fmt.Println("No Solution")
}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
    SimpleEquations(6, 6, 14) // 1 2 3
}
