package main

import "fmt"

func checkType(i interface{}) {
	switch v := i.(type) { // Получение типа переменной
	case int:
		fmt.Println("Это целое число:", v)
	case string:
		fmt.Println("Это строка:", v)
	case bool:
		fmt.Println("Это булево значение:", v)
	default:
		fmt.Println("Неизвестный тип")

	}
}

func main() {
	checkType(21)
	checkType("Hello")
	checkType(true)
}
