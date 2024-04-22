package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Функция, использующая канал для сигнализации остановки
func processChannel(stopChannel chan bool) {
	for {
		select {
		case <-stopChannel:
			fmt.Println("Горутина остановлена через канал")
			return
		default:
			fmt.Println("Горутина работает через канал")
			time.Sleep(time.Second)
		}
	}
}

// Функция, использующая общую переменную для сигнализации остановки
func processVar(stopIt *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if *stopIt {
			fmt.Println("Горутина остановлена через общую переменную")
			return
		}
		fmt.Println("Горутина работает через общую переменную")
		time.Sleep(time.Second)
	}
}

// Функция, использующая context для управления жизненным циклом горутины
func processContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина остановлена через context")
			return
		default:
			fmt.Println("Горутина работает через context")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// Использование канала для остановки
	stopChannel := make(chan bool)
	go processChannel(stopChannel)
	time.Sleep(3 * time.Second)
	stopChannel <- true
	time.Sleep(time.Second)

	// Использование общих переменных для остановки
	var stopIt bool
	var wg sync.WaitGroup
	wg.Add(1)
	go processVar(&stopIt, &wg)
	time.Sleep(3 * time.Second)
	stopIt = true
	wg.Wait()

	// Использование context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go processContext(ctx)
	<-ctx.Done()
	time.Sleep(time.Second)
}
