package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Counter struct {
	value int32
}

func (c *Counter) Increment() {
	atomic.AddInt32(&c.value, 1) // Инкрементируем счетчик атомарно
}

func (c *Counter) Value() int32 {
	return atomic.LoadInt32(&c.value) // Загружаем значение счетчика атомарно
}

func main() {
	counter := &Counter{}

	// Запуск горутин, которые инкрементируют счетчик
	for i := 0; i < 1000; i++ {
		go counter.Increment()
	}

	time.Sleep(time.Second)
	fmt.Println(counter.Value())
}
