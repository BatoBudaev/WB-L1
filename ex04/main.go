package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

// функция воркер
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Воркер %d выполняет %v\n", id, j)
		results <- j * j
	}
}

func main() {
	// чтение из stdin количества воркеров
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите количество воркеров: ")
	numWorkersStr, _ := reader.ReadString('\n')
	numWorkersStr = strings.TrimSpace(numWorkersStr)
	numWorkers, _ := strconv.Atoi(numWorkersStr)

	numJobs := 10
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup
	wg.Add(numWorkers) // увеличиваем счетчик WaitGroup на количество воркеров

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results, &wg)
	}

	// отправляем задания воркерам
	go func() {
		defer close(jobs) // закрываем канал jobs после отправки всех заданий
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
	}()

	// ожидание завершения работы всех воркеров
	go func() {
		wg.Wait()
		close(results) // закрываем канал результатов после завершения всех воркеров
	}()

	// обработка сигнала Ctrl+C
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("Ожидание сигнала")
	<-done
	// печать всех значений после выполнения
	for n := 1; n <= numJobs; n++ {
		fmt.Println(<-results)
	}
}
