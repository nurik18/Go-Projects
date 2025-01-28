package main

import (
	"fmt"
)

func main() {
	var inputNumber int
	fmt.Print("Введите число: ")
	if _, err := fmt.Scan(&inputNumber); err != nil {
		fmt.Println("Ошибка: введите целое число.")
		return
	}

	parity := "нечетное"
	if inputNumber%2 == 0 {
		parity = "четное"
	}
	fmt.Printf("Число %d - %s.\n", inputNumber, parity)
}
