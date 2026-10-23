package main

import "fmt"

func main() {

	var a, b float64
	var operation string

	for {
		fmt.Print("Введите первое число: ")

		_, err := fmt.Scan(&a)

		if err == nil {
			break
		}

		fmt.Println("Ошибка! Введите число.")
	}

	for {
		fmt.Print("Введите второе число: ")

		_, err := fmt.Scan(&b)

		if err == nil {
			break
		}

		fmt.Println("Ошибка! Введите число.")
	}

	fmt.Print("Введите операцию (+, -, *, /): ")
	_, err := fmt.Scan(&operation)

	if err != nil {
		fmt.Println("Ошибка: не удалось прочитать операцию")
		return
	}

	switch operation {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b != 0 {
			fmt.Println(a / b)
		} else {
			fmt.Println("Ошибка: деление на ноль")
		}
	default:
		fmt.Println("Неизвестная операция")
	}
}
