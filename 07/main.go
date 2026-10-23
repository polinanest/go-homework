package main

import "fmt"

func main() {
	var n int
	var sum int

	fmt.Print("Введите число N: ")
	_, err := fmt.Scan(&n)

	if err != nil || n <= 0 {
		fmt.Println("Ошибка: введите целое число больше 0")
		return
	}

	for i := 1; i <= n; i++ {
		if i%2 == 0 && i%3 == 0 {
			sum += i
		}
	}

	fmt.Println("Сумма подходящих чисел:", sum)
}
