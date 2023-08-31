package main

import (
	"fmt"
)

// Definisikan interface untuk menghitung jarak perkiraan
type DistanceCalculator interface {
	CalculateDistance() float64
}

// Definisikan struct Car
type Car struct {
	carType string
	fuelin  float64
}

// Implementasikan method CalculateDistance() untuk Car
func (c Car) CalculateDistance() float64 {
	fuelEfficiency := 1.5 // L/mill
	return c.fuelin * fuelEfficiency
}

// Fungsi untuk menghitung dan mencetak hasil jarak perkiraan
func printEstimatedDistance(dc DistanceCalculator) {
	fmt.Printf("Car type: %s, Est. distance: %.2f\n", dc.(*Car).carType, dc.CalculateDistance())
}

func main() {
	// Inisialisasi objek Car
	car := Car{
		carType: "SUV",
		fuelin:  10.5,
	}

	// Memanggil fungsi printEstimatedDistance dengan parameter car
	printEstimatedDistance(&car)
}
