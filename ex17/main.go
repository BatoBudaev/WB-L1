package main

import "fmt"

func search(nums []int, target int) int {
	// left и right индексы начала и конца
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2 // Индекс середины

		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid // Если значение в середине, возвращаем mid
		}
	}

	return -1 // Если ничего не найдено, то возвращаем -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	index := search(nums, 9)

	fmt.Println(index)
}
