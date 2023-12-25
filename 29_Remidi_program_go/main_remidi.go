package main

import (
	"fmt"
	"strings"
)

func main() {
	garis()

	var string1 string // deklarasi string kosong
	string1 = "Nama Saya"
	string2 := "Rois" // deklarasi string pake literal
	fmt.Println(string1, string2)

	garis()

	// manipulasi string
	string1 = "Nama Kamu"
	string2 = "Yamaha"
	result := string1 + " " + string2 // menggabung string
	fmt.Println(result)

	garis()

	// mengetahui panjang string
	lengthString1 := len(string1)
	lengthString2 := len(string2)
	fmt.Print("Panjang String 1 = ")
	fmt.Println(lengthString1)
	fmt.Print("Panjang String 2 = ")
	fmt.Println(lengthString2)

	garis()

	// Akses Karakter
	char := string1[2]
	fmt.Println(char)

	garis()

	// perulangan pada string
	for index, char := range result{
		fmt.Print(index)
		fmt.Println(char)
	}

	garis()

	// manipulasi string dengan paket `string`
	lovercaseString1 := strings.ToLower(string1)
	uppercaseString2 := strings.ToUpper(string2)
	fmt.Println(lovercaseString1, uppercaseString2)

	garis()

	
}

func garis(){
	fmt.Println("----------------")
}