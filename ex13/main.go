package main

import "fmt"

func main() {
	a, b := 1, 2
	a, b = b, a // Меняем 2 числа местами

	fmt.Printf("a = %v, b = %v", a, b)
}
