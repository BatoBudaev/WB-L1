package main

import "fmt"

func setBit(num int64, i int, bit int) int64 {
	if bit == 0 {
		mask := ^(1 << i)        // Создаем маску, где все биты, кроме i-го, установлены в 1
		return num & int64(mask) // Применяем маску к числу с помощью операции И
	} else {
		mask := 1 << i           // Создаем маску, где только i-й бит установлен в 1
		return num | int64(mask) // Применяем маску к числу с помощью операции ИЛИ
	}
}

func main() {
	num := int64(100)
	i := 0
	bit := 1

	res := setBit(num, i, bit)
	fmt.Printf("Исходное число: %b, Измененное число: %b\n", num, res)
}
