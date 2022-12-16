package main

import (
	"fmt"
	"math"
)

func area(raio float64) float64 {
	return math.Pow(raio, 2) * 3.14159
}

func main() {
	var raio float64
	_, err := fmt.Scanf("%f", &raio)

	if err != nil {
		return
	}

	area := area(raio)
	result := fmt.Sprintf("%.4f", area)

	fmt.Println("A=" + result)
}
