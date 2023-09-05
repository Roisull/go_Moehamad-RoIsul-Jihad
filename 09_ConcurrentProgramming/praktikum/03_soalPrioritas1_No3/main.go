package main

import (
	"fmt"
	"time"
)

func printMultiplesOfThree(ch chan<- int) {
	for i := 3; ; i += 3 {
		ch <- i // Kirim angka kelipatan 3 ke channel.
		time.Sleep(3 * time.Second)
	}
}

func main() {
	bufferSize := 5 // Ukuran buffer channel.
	ch := make(chan int, bufferSize)

	go printMultiplesOfThree(ch)

	for i := 0; i < 5; i++ { // Mencetak 5 angka kelipatan 3.
		multiple := <-ch
		fmt.Printf("Kelipatan 3: %d\n", multiple)
	}

	close(ch) // Menutup channel setelah selesai.

	fmt.Println("Program selesai.")
}