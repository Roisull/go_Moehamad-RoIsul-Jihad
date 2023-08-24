package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan nilai n: ")
    fmt.Scanln(&n)

    for i := 1; i <= n; i++ {
        // Mencetak spasi sebelum asterisk
        for j := 1; j <= n-i; j++ {
            fmt.Print(" ")
        }

        // Mencetak asterisk
        for j := 1; j <= i; j++ {
            fmt.Print("* ")
        }

        // Pindah baris setelah setiap baris selesai dicetak
        fmt.Println()
    }
}