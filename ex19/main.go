package main

import "fmt"

func reverseString(str string) string {
	runes := []rune(str)
	length := len(runes)

	// Переворачиваем массив рун
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Возвращаем перевернутую строку
	return string(runes)
}

func main() {
	str := "главрыба — абырвалг"
	reversed := reverseString(str)
	fmt.Println(reversed)
}
