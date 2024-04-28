package main

import "fmt"

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	stack := [][]int{arr} // Инициализируем стек, помещаем в него исходный массив
	var sorted []int      // Массив для хранения отсортированных значений

	for len(stack) > 0 {
		current := stack[len(stack)-1] // Получаем верхний элемент из стека
		stack = stack[:len(stack)-1]

		// Если длина текущего массива меньше или равна 1, добавляем его
		// в отсортированный массив и переходим к следующему итерации
		if len(current) <= 1 {
			sorted = append(sorted, current...)
			continue
		}

		pivot := current[len(current)/2] // Опорный элемент
		left := make([]int, 0)           // Меньше опорного
		right := make([]int, 0)          // Больше опорного

		for _, num := range current {
			if num < pivot {
				left = append(left, num)
			} else if num > pivot {
				right = append(right, num)
			}
		}

		stack = append(stack, right, []int{pivot}, left) // Помещаем right, pivot и left в стек для дальнейшей обработки
	}

	return sorted // Возвращаем отсортированный массив
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Исходный массив:", arr)
	sortedArr := QuickSort(arr)
	fmt.Println("Отсортированный массив:", sortedArr)
}
