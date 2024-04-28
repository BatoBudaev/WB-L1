package main

import (
	"fmt"
	"time"
)

// Собственная функция sleep
func Sleep(t time.Duration) {
	ch := time.After(t) // Создаем канал, который будет закрыт после истечения времени t
	for {
		select {
		case <-ch: // Если канал закрыт, то завершаем функцию
			return
		}
	}
}

func main() {
	fmt.Println("Перед сном")
	Sleep(3 * time.Second) // Пауза на 3 секунды
	fmt.Println("3 секунды прошли")
}
