package main

import (
	"errors"
	"fmt"
)

func selectSort(arr []int) ([]int, error) {
	arrLength := len(arr)
	if arrLength == 0 {
		return arr, errors.New("Массив пустой")
	}

	for i := 0; i < arrLength; i++ {
		minIndex := i
		for j := i + 1; j < arrLength; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr, nil
}

func main() {
	arr := []int{5, 22, 3, 1, 23, 0}
	arr, err := selectSort(arr)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(arr)
}
