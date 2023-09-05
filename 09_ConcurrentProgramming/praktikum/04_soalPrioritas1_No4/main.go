package main

import (
    "fmt"
    "sync"
)

// Fungsi untuk menghitung faktorial dari suatu bilangan.
func calculateFactorial(n int, wg *sync.WaitGroup, mu *sync.Mutex) {

    defer wg.Done()
    result := 1

    for i := 1; i <= n; i++ {
        // Mengunci mutex sebelum mengubah variabel yang dibagikan.
        mu.Lock()
        result *= i
        // Membuka kunci mutex setelah selesai mengubah variabel.
        mu.Unlock()
    }
    fmt.Printf("Faktorial dari %d adalah %d\n", n, result)
}

func main() {
    num := 5 // Ganti bilangan untuk dihitung faktorialnya.
    var wg sync.WaitGroup
    var mu sync.Mutex

    wg.Add(1)
    go calculateFactorial(num, &wg, &mu)

    wg.Wait()
}