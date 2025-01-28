package main

import (
	"fmt"
)

func main() {
	fmt.Println("Вас приветсвует программа реверса строки")
	fmt.Print("Введите ваше слово: ")
	var inputString string
	if _, err := fmt.Scan(&inputString); err != nil || inputString == "" {
		fmt.Println("Ошибка: введите непустую строку.")
		return
	}
	fmt.Print("Реверс: ")
	fmt.Println(stringReverse(inputString))
}

func stringReverse(word string) string {
	wordRune := []rune(word)
	n := len(wordRune)
	newRune := make([]rune, n)
	for i, r := range wordRune {
		newRune[n-1-i] = r
	}
	return string(newRune)
}
