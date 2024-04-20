package main

import (
	"fmt"
	"sync"
)

// функция вычисления квадрата
func square(n int) int {
	return n * n
}

func main() {
	a := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, n := range a {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			fmt.Println(square(val))
		}(n) // Передаём n как аргумент
	}

	wg.Wait() // Ждём пока все горутины закончатся
}
