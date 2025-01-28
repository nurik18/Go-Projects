package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Description string
	Done        bool
}

func main() {
	menu()
}

func menu() {
	var tasks []Task
	for {
		fmt.Println("===== Менеджер задач =====")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Просмотреть список")
		fmt.Println("3. Отметить выполненной")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Выйти")
		fmt.Println()
		fmt.Print("Выберите действие: ")
		var t int
		fmt.Scanln(&t)
		fmt.Println()
		switch t {
		case 1:
			fmt.Print("Введите описание задачи: ")
			reader := bufio.NewReader(os.Stdin) // Создаём буферизированный ридер
			input, _ := reader.ReadString('\n') // Считываем строку с пробелами
			input = strings.TrimSpace(input)    // Убираем лишний перевод строки
			tasks = append(tasks, Task{input, false})
			fmt.Printf("Задача добавлена! \n\n")
		case 2:
			for i, value := range tasks {
				done := "❌"
				if value.Done {
					done = "✅"
				}
				fmt.Printf("%d. %s - %s\n", i+1, value.Description, done)
			}
			fmt.Println()
		case 3:
			for i, value := range tasks {
				done := "❌"
				if value.Done {
					done = "✅"
				} else {
					fmt.Printf("%d. %s - %s\n", i+1, value.Description, done)
				}
			}
			fmt.Print("Выберите задачу для отметки: ")
			var numb int
			fmt.Scanln(&numb)
			tasks[numb-1].Done = true
			fmt.Println("Задача отмечена!")
			fmt.Println()
		case 4:
			for i, value := range tasks {
				done := "❌"
				if value.Done {
					done = "✅"
				}
				fmt.Printf("%d. %s - %s\n", i+1, value.Description, done)
			}
			fmt.Print("Выберите задачу для удаления: ")
			var numb int
			fmt.Scanln(&numb)
			tasks = append(tasks[:numb-1], tasks[numb:]...)
			fmt.Println("Задача удалена!")
			fmt.Println()
		case 5:
			fmt.Println("Выход...")
			return
		default:
			fmt.Println("Некорректный ввод! Попробуйте снова.")
		}
	}
}
