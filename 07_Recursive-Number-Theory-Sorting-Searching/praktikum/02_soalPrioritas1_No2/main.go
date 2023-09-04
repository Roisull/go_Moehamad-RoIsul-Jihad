package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	counts := make(map[string]int)

	// Menghitung jumlah kemunculan setiap item
	for _, item := range items {
		counts[item]++
	}

	// Membuat slice dari pasangan item dan jumlah kemunculannya
	var pairs []pair
	for name, count := range counts {
		pairs = append(pairs, pair{name, count})
	}

	// Mengurutkan slice berdasarkan jumlah kemunculan
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count < pairs[j].count
	})

	return pairs
}

func main() {
	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"})
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"football", "basketball", "tennis"})
	for _, list := range pairs {
		fmt.Print(list.name, " -> ", list.count, " ")
	}
}
