package main

import "fmt"

func main() {
	values := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]bool) // Map для хранения множества

	// Добавление элементов в множество
	for _, item := range values {
		set[item] = true
	}

	fmt.Println(set)
}
