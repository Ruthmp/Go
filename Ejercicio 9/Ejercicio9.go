package main

import (
	"fmt"
)

func main() {
	//Con este programa, lee todo el String, pero si se introduce por teclado solamente lee la primera palabra

	var myString = "Esto es una cadena de texto"
	fmt.Println("Introduce texto")
	fmt.Scanln(&myString)

	contador := 0

	for _, c := range myString {
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			contador++
		}
	}
	fmt.Println("Número de vocales: ", contador)
}
