package main

import "fmt"

func main() {
	var n int
	fmt.Println("Вас приветствует программа сортировки массива")
	var arr []int
	fmt.Print("Введите кол-во элементов:")
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Printf("Ошибка: %s\n", err)
		return
	}
	fmt.Println("Введите ваши элементы:")
	for i := 0; i < n; i++ {
		var t int
		if _, err := fmt.Scan(&t); err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			return
		}
		arr = append(arr, t)
	}
	bubbleSort(&arr)
	fmt.Print("Ваш массив: ")
	fmt.Println(arr)
}

func bubbleSort(arr *[]int) {
	n := len(*arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}
