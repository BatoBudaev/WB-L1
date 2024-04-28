package main

import "fmt"

func arithmeticOperations(a int64, b int64) {
	// Сложение
	sum := a + b
	fmt.Printf("Сложение: %d + %d = %d\n", a, b, sum)

	// Вычитание
	sub := a - b
	fmt.Printf("Вычитание: %d - %d = %d\n", a, b, sub)

	// Перемножение
	product := a * b
	fmt.Printf("Перемножение: %d * %d = %d\n", a, b, product)

	// Деление
	div := a / b
	fmt.Printf("Деление: %d / %d = %d\n", a, b, div)

}

func main() {
	a := int64(1 << 21) // 2^21
	b := int64(1 << 21) // 2^21

	arithmeticOperations(a, b)
}
