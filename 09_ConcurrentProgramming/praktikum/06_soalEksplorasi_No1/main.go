package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Product struct {
	ID		int64 		`json:"id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func aksesProduk(url string, startID int, count int, wg *sync.WaitGroup, ch chan<- Product) {
	defer wg.Done()

	resp, err := http.Get(fmt.Sprintf("%s?start=%d&limit=%d", url, startID, count))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var products []Product
	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, product := range products {
		ch <- product
	}
}

func main() {
	url := "https://fakestoreapi.com/products"
	startID := 1
	count := 5

	var wg sync.WaitGroup
	channelProduk := make(chan Product, count) // Buffer channel untuk menampung hasil

	wg.Add(1)
	go aksesProduk(url, startID, count, &wg, channelProduk)

	// for i := 0; i < numWorkers; i++ {
	// 	wg.Add(1)
	// 	go fetchProducts(url, &wg, productsChannel)
	// }

	go func() {
		wg.Wait()
		close(channelProduk) // Menutup channel setelah semua goroutine selesai
	}()

	fmt.Println("Products data")
	fmt.Println("===")

	for product := range channelProduk {
		fmt.Println("ID:", product.ID)
		fmt.Println("title:", product.Title)
		fmt.Println("price:", product.Price)
		fmt.Println("category:", product.Category)
		fmt.Println("===")
	}
}