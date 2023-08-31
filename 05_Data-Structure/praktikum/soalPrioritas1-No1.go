package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	merged := make(map[string]bool) // Membuat map untuk memastikan nilai unik
	result := []string{}

	for _, val := range arrayA {
		if !merged[val] {
			merged[val] = true
			result = append(result, val)
		}
	}

	for _, val := range arrayB {
		if !merged[val] {
			merged[val] = true
			result = append(result, val)
		}
	}

	return result
}

func soalArray() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]

	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}