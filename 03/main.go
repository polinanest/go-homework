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
		sum += i
	}

	average := float64(sum) / float64(n)

	fmt.Println("Сумма чисел:", sum)
	fmt.Println("Среднее значение:", average)
}
