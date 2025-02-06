package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("notes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}

	defer file.Close()

	fmt.Print("Введите текст для сохранения: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	_, err = file.WriteString(input + "\n")
	if err != nil {
		fmt.Println("Ошибка при записи в файл: ", err)
		return
	}
	fmt.Println("Запись сохранена в notes.txt!")

	fmt.Println("\nСодержимое notes.txt:")
	readFile, err := os.Open("notes.txt")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	defer readFile.Close()

	reader := bufio.NewScanner(readFile)
	for reader.Scan() {
		fmt.Println(reader.Text())
	}

	if err := reader.Err(); err != nil {
		fmt.Println("Ошибка при чтении:", err)
	}
}
