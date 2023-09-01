package main

import "fmt"

func caesar(offset int, input string) string {
    result := ""
    for _, char := range input {
        if char == ' ' {
            result += " "
        } else {
            shiftedChar := 'a' + (char-'a'+rune(offset))%26 // Menggunakan operator modulo untuk menghindari pergeseran melewati 'z'
            result += string(shiftedChar)
        }
    }
    return result
}

func main() {
    fmt.Println(caesar(3, "abc"))                  // def
    fmt.Println(caesar(2, "alta"))                // cnav
    fmt.Println(caesar(10, "alterraacademy"))     // kvvdxxkkksmfli
    fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza
    fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
