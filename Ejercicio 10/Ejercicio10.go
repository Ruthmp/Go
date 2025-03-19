package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generarContrasena(longitud int) string {
	validChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	//Crear una variable para almacenar la contraseña
	var password string

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < longitud; i++ {
		indice := rand.Intn(len(validChars))
		password += string(validChars[indice])
	}
	// Devolver la contraseña generada
	return password
}

func main() {
	longitud := 10
	password := generarContrasena(longitud)
	fmt.Println("Contraseña generada: ", password)
}
