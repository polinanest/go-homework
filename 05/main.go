package main

import "fmt"

func main() {
	var n int
	var number float64
	var min, max float64

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

		if i == 0 {
			min = number
			max = number
		} else {
			if number < min {
				min = number
			}

			if number > max {
				max = number
			}
		}
	}

	fmt.Println("Минимальное число:", min)
	fmt.Println("Максимальное число:", max)
}
