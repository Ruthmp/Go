package main

import "fmt"

func esPrimo(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	numero := 5
	if esPrimo(numero) {
		fmt.Println(numero, " es primo.")
	} else {
		fmt.Println(numero, "no es primo.")
	}

}
