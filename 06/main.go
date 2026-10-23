package main

import "fmt"

func main() {
	var n int
	var number float64
	var positive, negative, zeros int

	fmt.Print("Введите количество чисел: ")
	_, err := fmt.Scan(&n)

	if err != nil || n <= 0 {
		fmt.Println("Ошибка: введите целое число больше 0")
		return
	}

	fmt.Printf("Введите %d чисел через пробел: ", n)

	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&number)

		if err != nil {
			fmt.Println("Ошибка: необходимо ввести число")
			return
		}

		if number > 0 {
			positive++
		} else if number < 0 {
			negative++
		} else {
			zeros++
		}
	}

	fmt.Println("Положительных чисел:", positive)
	fmt.Println("Отрицательных чисел:", negative)
	fmt.Println("Нулей:", zeros)
}
