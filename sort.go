package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 3, 10, 9, 4, 5, 8, 2, 7, 6, 0, -1}
	fmt.Printf("Начальный массив: \n%v\n", arr)
	bubbleSort(arr)
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Printf("Отсортированный массив: \n%v \n", arr)
}
