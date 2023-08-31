package main

import (
    "fmt"
    "strings"
)
func palindromeChecker() {
    var input string
    fmt.Println("Apakah palindrome?")
    fmt.Print("Masukkan kata: ")

    // Menggunakan Scanner untuk menangkap input dari konsol
    fmt.Scanln(&input)
    input = strings.ToLower(input) // Mengubah input menjadi huruf kecil
    // Menghapus spasi dari input
    input = strings.ReplaceAll(input, " ", "")
    // Memeriksa apakah input adalah palindrome atau bukan
    isPalindrome := true
    for i := 0; i < len(input)/2; i++ {
        if input[i] != input[len(input)-1-i] {
            isPalindrome = false
            break
        }
    }
    // Menampilkan hasil
    if isPalindrome {
        fmt.Println("Captured:", input)
        fmt.Println("Palindrome")
    } else {
        fmt.Println("Captured:", input)
        fmt.Println("Bukan palindrome")
    }
}
