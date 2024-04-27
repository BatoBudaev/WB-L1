package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cs := map[string]int{"касса": 0}
	mu := sync.Mutex{} // Определяем мьютекс

	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()         // Блокируем доступ к map
			defer mu.Unlock() // Через defer разблокируем

			cs["касса"] += 1
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(cs)
}
