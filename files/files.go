package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("файл не найден: %v", err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %v", err)
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return fmt.Errorf("ошибка при копировании файла: %v", err)
	}
	return nil
}

func main() {
	err := copyFile("source.txt", "copy.txt")
	if err != nil {
		fmt.Println("ошибка: ", err)
	} else {
		fmt.Println("Файл успешно скопирован!")
	}
}
