package main

import "fmt"

//todo переписать на хэш таблицы
func fibonacci(n uint) uint {
	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}
func main() {
	// todo добавить ввод с консоли
	fmt.Println(fibonacci(10))
}
