//Leer un archivo `.txt` e imprimir su contenido.

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Crear archivo
	file, err := os.Create("archivo.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo", err)
		return
	}
	defer file.Close()

	//Escribir archivo
	_, err = file.WriteString("Esto es una prueba")
	if err != nil {
		fmt.Println("Error de escritura", err)
		return
	}

	//Abrir el archivo

	file, err = os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo", err)
		return
	}
	defer file.Close()

	//Leer archivo
	reader, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error de lectura", err)
		return
	}
	//Mostrar contenido
	fmt.Println("Contenido del archivo: \n", string(reader))

}
