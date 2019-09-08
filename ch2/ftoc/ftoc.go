// Ftoc exibe duas conversões de Fahrenheit para Celsius
package main

import "fmt"

func main() {
	const freezingF, bolingF = 32.0, 212.0
	fmt.Printf("Freezing point = %g°F or %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("Boiling point = %g°F or %g°C\n", bolingF, fToC(bolingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
