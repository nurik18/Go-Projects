package main

import "fmt"

type Expense struct {
	Category string
	Amount   float64
	Date     string
}

func main() {
	fmt.Println("ВНИМАНИЕ !!! Программа не защищена от ошибок пользователя. (Это не трудно, просто лень)")
	fmt.Println()
	var n int
	var expenses []Expense
	for {
		menu()
		n = act()
		switch n {
		case 1:
			act1(&expenses)
		case 2:
			act2(&expenses)
		case 3:
			act3(&expenses)
		case 4:
			act4(&expenses)
		case 5:
			act5(&expenses)
		case 6:
			fmt.Println("Выход...")
			return
		default:
			fmt.Println("Некорректный ввод! Попробуйте снова.")
			fmt.Println()
		}
	}
}

func menu() {
	fmt.Println()
	fmt.Println("===== Трекер расходов =====")
	fmt.Println("1. Добавить расход")
	fmt.Println("2. Просмотреть все расходы")
	fmt.Println("3. Фильтровать по категории")
	fmt.Println("4. Удалить расход")
	fmt.Println("5. Общая сумма расходов")
	fmt.Println("6. Выйти")
	fmt.Println()
}

func act() int {
	var n int
	fmt.Print("Выберите действие: ")
	fmt.Scanln(&n)
	return n
}

func act1(expenses *[]Expense) {
	var s string
	var sum float64
	var date string
	fmt.Print("Введите категорию: ")
	fmt.Scanln(&s)
	fmt.Print("Введите сумму: ")
	fmt.Scanln(&sum)
	fmt.Print("Введите дату: ")
	fmt.Scanln(&date)
	*expenses = append(*expenses, Expense{s, sum, date})
}

func act2(expenses *[]Expense) {
	fmt.Println("№   Категория         Сумма      Дата")
	for i, value := range *expenses {
		fmt.Printf("%-3d %-17s %-10.2f %s\n", i+1, value.Category, value.Amount, value.Date)
	}
	fmt.Println()
}

func act3(expenses *[]Expense) {
	fmt.Println("№   Категория")
	for i, value := range *expenses {
		fmt.Printf("%-3d   %s\n", i+1, value.Category)
	}
	var n int
	fmt.Print("Обозначьте номер категории, которой хотите вывести: ")
	fmt.Scanln(&n)
	s := (*expenses)[n-1].Category
	var filteredExpenses []Expense
	for _, value := range *expenses {
		if value.Category == s {
			filteredExpenses = append(filteredExpenses, value)
		}
	}
	fmt.Println("№   Сумма      Дата")
	for i, value := range filteredExpenses {
		fmt.Printf("%-3d %-10.2f %s\n", i+1, value.Amount, value.Date)
	}
}

func act4(expenses *[]Expense) {
	fmt.Println("№   Категория         Сумма      Дата")
	for i, value := range *expenses {
		fmt.Printf("%-3d %-17s %-10.2f %s\n", i+1, value.Category, value.Amount, value.Date)
	}
	fmt.Println()
	var n int
	fmt.Print("Обозначьте номер категории, которой хотите удалить: ")
	fmt.Scanln(&n)
	*expenses = append((*expenses)[:n-1], (*expenses)[n:]...)
	fmt.Println("Категория удалена!")
}

func act5(expenses *[]Expense) {
	var sum float64
	for _, value := range *expenses {
		sum += value.Amount
	}
	fmt.Printf("Общая сумма ваших расходов: %-10.2f", sum)
	fmt.Println()
}
