package main

import "fmt"

func main() {
	var n int
	var number int
	var sum int
	var min, max int
	var even, odd int

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
			fmt.Println("Ошибка: необходимо ввести целое число")
			return
		}

		sum += number

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

		if number%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	average := float64(sum) / float64(n)

	fmt.Println("\nСтатистика:")
	fmt.Println("Сумма:", sum)
	fmt.Printf("Среднее: %.2f\n", average)
	fmt.Println("Максимум:", max)
	fmt.Println("Минимум:", min)
	fmt.Println("Чётных чисел:", even)
	fmt.Println("Нечётных чисел:", odd)
}
