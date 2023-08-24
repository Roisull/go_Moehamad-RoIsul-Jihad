package main

import "fmt"

func main() {
    var num int
    fmt.Print("Masukkan angka: ")
    fmt.Scanln(&num)

    fmt.Printf("Faktor dari %d adalah: ", num)
    for i := 1; i <= num; i++ {
        if num%i == 0 {
            fmt.Printf("%d ", i)
        }
    }
    fmt.Println()
}