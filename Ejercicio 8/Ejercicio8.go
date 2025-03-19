package main

import "fmt"

func main() {
	var celsius float64
	fmt.Println("Introduce la temperatura en Celsius")
	fmt.Scanln(&celsius)

	// Mostrar el valor ingresado
	fmt.Println("La temperatura en Celsius es:", celsius)

	// Pasar Celsius a Fahrenheit

	var fahrenheit float64
	fahrenheit = ((celsius * 9) / 5) + 32

	fmt.Println("La temperatura en Fahrenheit es:", fahrenheit)
}
