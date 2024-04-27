package main

import "fmt"

func findIntersection(set1, set2 []int) []int {
	m := make(map[int]bool) // Map для отслеживания элементов первого слайса
	for _, item := range set1 {
		m[item] = true
	}

	var intersection []int // Cлайс для хранения пересечения
	for _, item := range set2 {
		if _, ok := m[item]; ok {
			intersection = append(intersection, item) // Если элемент присутствует в map, добавляем его в пересечение
		}
	}

	return intersection
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{5, 4, 6, 7, 3}

	intersection := findIntersection(set1, set2)
	fmt.Println("Пересечение:", intersection)

}
