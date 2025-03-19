package main

import "fmt"

func main() {
	for num := 1; num <= 10; num++ {
		if num%2 == 0 {
			fmt.Println(num)
		}
	}
}
