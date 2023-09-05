package main

import (
    "fmt"
	"sync"
    "time"
)

// Fungsi untuk mencetak angka kelipatan x dengan interval waktu 3 detik.
func printMultiples(x int, wg *sync.WaitGroup) {
	defer wg.Done()
    for i := 1; i <= 10; i++ { // Menghentikan goroutine setelah mencetak 10 kelipatan.
		fmt.Printf("Kelipatan %d: %d\n", x, x*i)
		time.Sleep(3 * time.Second)
	}
}

func main() {
    x := 5 // Ganti nilai x sesuai dengan kelipatan yang diinginkan.

    var wg sync.WaitGroup
	wg.Add(1)
	go printMultiples(x, &wg)

	wg.Wait() // menunggu  goroutine selesai
	fmt.Println("Program Selesai")
}