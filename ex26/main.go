package main

import (
	"fmt"
	"strings"
)

func unique(s string) bool {
	s = strings.ToLower(s) // Преобразование в нижний регистр

	m := make(map[rune]bool) // Использование map для отслеживания уникальности символов
	for _, char := range s {
		if m[char] {
			return false // Если символ уже встречался, возвращаем false
		}
		m[char] = true
	}

	return true // Если все символы уникальны, возвращаем true
}

func main() {
	fmt.Println(unique("abcd"))      // true
	fmt.Println(unique("abCdefAaf")) // false
	fmt.Println(unique("aabcd"))     // false
}
