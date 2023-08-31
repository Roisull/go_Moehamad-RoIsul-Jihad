package praktikum

import (
	"fmt"
)

func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	
	if number <= 3 {
		return true
	}
	
	if number%2 == 0 || number%3 == 0 {
		return false
	}
	
	for i := 5; i*i <= number; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}
	
	return true
}

func primeNumber(number int) string {
	if isPrime(number) {
		return "Bilangan Prima"
	}
	return "Bukan Bilangan Prima"
}

func main() {
	fmt.Println(primeNumber(1000000007)) // Bilangan Prima
	fmt.Println(primeNumber(13))         // Bilangan Prima
	fmt.Println(primeNumber(17))         // Bilangan Prima
	fmt.Println(primeNumber(20))         // Bukan Bilangan Prima
	fmt.Println(primeNumber(35))         // Bukan Bilangan Prima
}
