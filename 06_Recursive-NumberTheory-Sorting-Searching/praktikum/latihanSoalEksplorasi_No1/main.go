package main

import (
    "fmt"
    "strings"
)

type student struct {
    name       string
    // nameEncode string
    // score      int
}

type Chiper interface {
    Encode() string
    Decode() string
}

func (s *student) Encode() string {
    var nameEncode strings.Builder

    for _, char := range s.name {
        if char == ' ' {
            nameEncode.WriteRune(' ')
        } else {
            shiftedChar := 'a' + ((char-'a')+1)%26 // Menggunakan offset 1 untuk enkripsi
            nameEncode.WriteRune(shiftedChar)
        }
    }

    return nameEncode.String()
}

func (s *student) Decode() string {
    var nameDecode strings.Builder

    for _, char := range s.name {
        if char == ' ' {
            nameDecode.WriteRune(' ')
        } else {
            shiftedChar := 'a' + ((char-'a')+25)%26 // Menggunakan offset -1 untuk dekripsi
            nameDecode.WriteRune(shiftedChar)
        }
    }

    return nameDecode.String()
}

func main() {
    var menu int
    var a student
    var c Chiper = &a

    fmt.Print("[1] Encrypt\n[2] Decrypt\nChoose your menu? ")
    fmt.Scan(&menu)

    fmt.Print("\nInput Student Name: ")
    fmt.Scan(&a.name)

    if menu == 1 {
        fmt.Printf("\nEncode of student's name %s is : %s\n", a.name, c.Encode())
    } else if menu == 2 {
        fmt.Printf("\nDecode of student's name %s is : %s\n", a.name, c.Decode())
    }
}
