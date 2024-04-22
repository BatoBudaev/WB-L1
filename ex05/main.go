package main

import (
	"fmt"
	"time"
)

func main() {
	values := make(chan int, 1)
	timeout := time.After(5 * time.Second) // Задаем таймаут в 5 секунд

	// Запускаем горутину для отправки значений в канал
	go func() {
		defer close(values)
		for i := 0; i < 10; i++ {
			values <- i
			time.Sleep(1 * time.Second)
		}
	}()

	// Читаем значения из канала с таймаутом
	for {
		select {
		case value, ok := <-values:
			if !ok {
				fmt.Println("Канал закрыт")
				return
			}
			fmt.Println("Получено значение: ", value)
		case <-timeout:
			fmt.Println("Время вышло")
			return
		}
	}
}
