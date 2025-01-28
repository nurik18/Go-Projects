package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Вас приветствует программа проверки числа на простату")
	fmt.Print("Введите ваше число: ")
	var inputNumber int
	if _, err := fmt.Scan(&inputNumber); err != nil {
		fmt.Printf("Ошибка: %s\n", err)
		return
	}
	fmt.Print("Является ли ваше число простым? ")
	fmt.Println(isPrime(inputNumber))
}

func isPrime(n int) string {
	if n <= 0 {
		return "НЕТ"
	}
	if n == 1 || n == 2 {
		return "ДА"
	}
	if n%2 == 0 {
		return "НЕТ"
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i < sqrtN; i += 2 {
		if n%i == 0 {
			return "НЕТ"
		}
	}
	return "ДА"
}
