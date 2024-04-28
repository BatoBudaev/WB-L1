package main

import "fmt"

// Функция для удаления элемента по индексу из слайса
func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	s := []int{10, 20, 30, 40, 50, 60}
	fmt.Println("Исходный слайс:", s)

	index := 3 // Индекс элемента для удаления
	s = removeElement(s, index)
	fmt.Println("Слайс после удаления:", s)
}
