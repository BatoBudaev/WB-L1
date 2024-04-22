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
	res := make(chan int) // канал для передачи вычислений квадратов
	var wg sync.WaitGroup // для ожидания завершения всех горутин

	for _, n := range a {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			res <- square(val) // отправка в канал
		}(n)
	}

	go func() {
		wg.Wait()
		close(res) // закрытие канала
	}()

	sum := 0
	for v := range res {
		sum += v // подсчет всей суммы
	}

	fmt.Println(sum)
}
