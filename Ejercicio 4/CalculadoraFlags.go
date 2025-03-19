package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Definición de los flags
	var suma bool
	var resta bool

	flag.BoolVar(&suma, "a", false, "Operación de suma")
	flag.BoolVar(&resta, "b", false, "Operación de resta")

	// Parseamos los flags
	flag.Parse()

	// Verificamos que se hayan pasado exactamente dos números
	if len(flag.Args()) != 2 {
		fmt.Println("Por favor, proporciona dos números para operar.")
		os.Exit(1)
	}

	// Convertir los argumentos de la línea de comandos a enteros
	var num1, num2 int
	_, err := fmt.Sscanf(flag.Args()[0], "%d", &num1)
	if err != nil {
		fmt.Println("Error: El primer número no es válido.")
		os.Exit(1)
	}
	_, err = fmt.Sscanf(flag.Args()[1], "%d", &num2)
	if err != nil {
		fmt.Println("Error: El segundo número no es válido.")
		os.Exit(1)
	}

	// Realizar la operación basada en el flag
	if suma {
		resultado := num1 + num2
		fmt.Printf("La suma de %d y %d es: %d\n", num1, num2, resultado)
	} else if resta {
		resultado := num1 - num2
		fmt.Printf("La resta de %d y %d es: %d\n", num1, num2, resultado)
	} else {
		fmt.Println("Por favor, selecciona una operación válida (-a para suma, -b para resta).")
		os.Exit(1)
	}

}
