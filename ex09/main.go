package main

import "fmt"

func generate(nums ...int) <-chan int {
	res := make(chan int)
	go func() {
		for _, n := range nums { // Чтение из массива
			res <- n // Отправка чисел в канал
		}
		close(res)
	}()

	return res
}

func double(ch <-chan int) <-chan int {
	res := make(chan int)
	go func() {
		for n := range ch { // Чтение из канала
			res <- n * 2 // Отправка чисел умноженных на 2 в канал
		}
		close(res)
	}()

	return res
}

func main() {
	ch := generate(1, 2, 3, 4)
	res := double(ch)

	for n := range res {
		fmt.Println(n) // Печать из канала
	}
}
