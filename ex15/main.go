package main

import "fmt"

// Проблема была в том, что justString = v[:100] ссылается на исходный массив v. v остается в памяти, даже если используется только его часть justString.

func createHugeString(size int) string {
	return string(make([]byte, size))
}

func someFunc() {
	v := createHugeString(1 << 10)  // Создание большой строки
	justString := make([]byte, 100) // Создание среза байтов для новой строки
	copy(justString, v[:100])       // Копирование первых 100 байтов из v в justString
	fmt.Println(string(justString)) // Преобразование среза байтов в строк
}

func main() {
	someFunc()
}
