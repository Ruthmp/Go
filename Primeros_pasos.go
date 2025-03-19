package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {

	//Hola mundo

	fmt.Println(("Hola Go!"))

	/*
		Esto es un comentario
	*/

	//!Variables

	//String

	var myString string = "Esto es una cadena de texto"
	fmt.Println(myString)

	myString = "Aquí cambio el valor de la cadena de TXT"
	fmt.Println(myString)

	var myString2 = "Esto es una cadena de texto" //Podemos declarar la variable sin indicar el tipo de dato
	fmt.Println(myString2)

	//Int
	var myInt = 7
	fmt.Println(myInt)

	var myInt2 int = 7
	fmt.Println(myInt2)

	myInt = myInt + 4
	fmt.Println(myInt)
	fmt.Println(myInt - 1)
	fmt.Println(myInt)

	fmt.Println(myString, myInt)

	fmt.Println(reflect.TypeOf(myInt))

	//Float

	var myFloat float64 = 6.5

	fmt.Println((myFloat))
	fmt.Println(reflect.TypeOf(myFloat))

	fmt.Println(float64(myInt) + myFloat) //Para formatear myInt en float y no perder decimales

	//Boolean

	var myBool bool = false
	myBool = true
	fmt.Println(myBool)

	myString3 := "Esto es una cadena de TXT" // con := declara e inicializamos, sin necesidad de poner "var"
	fmt.Println(myString3)

	//constantes

	const myConst = "Esto es una constante"
	fmt.Println(myConst)

	//!Control de Flujo

	if myInt == 10 && myString == "Hola" {
		fmt.Println("El valor es 10")
	} else if myInt == 11 || myString == "Hola" {
		fmt.Println("El valor es 11")

	} else {
		fmt.Println("El valor no es 10")
	}

	//!Estructuras de datos

	//Array

	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	fmt.Println(myArray)
	fmt.Println(myArray[2])

	//Map

	myMap := make(map[string]int)
	myMap["Ruth"] = 35
	myMap["Zara"] = 27
	myMap["Penny"] = 10
	fmt.Println(myMap)
	fmt.Println(myMap["Ruth"])

	myMap2 := map[string]int{"Ruth": 35, "Zara": 27, "Penny": 10}
	fmt.Println(myMap2)

	//Lista

	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	fmt.Println(myList.Back().Value)

	//!Bucles

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}
	for index, value := range myArray {
		fmt.Println(index, value)
	}

	//!Funciones

	myFunction()           //Llamar a myFunction
	fmt.Println(myFunc2()) //Llamar a miFunc2

	//!Estructuras

	type MyStruct struct {
		name string
		age  int
	}

	myStruct := MyStruct{"Ruth", 35}
	fmt.Println(myStruct)
}
func myFunction() {
	fmt.Println("Mi función")
}

func myFunc2() string {
	return "Mi Función 2"
}
