package main

import "fmt"

func main() {
	fmt.Println("Вас приветсвует программа которая принимает число n и возвращает n-е число Фибоначчи")
	fmt.Print("Введите число: ")
	var inputNumber int
	fmt.Scan(&inputNumber)
	fib := fibonacci(inputNumber)
	fmt.Printf("Ваше число: %d\n", fib)
}

func fibonacci(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
