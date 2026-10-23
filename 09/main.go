package main

import "fmt"

func main() {
	var balance, amount float64
	var action int

	fmt.Print("Введите начальный баланс: ")
	_, err := fmt.Scan(&balance)

	if err != nil || balance < 0 {
		fmt.Println("Ошибка: некорректный баланс")
		return
	}

	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1 — Посмотреть баланс")
		fmt.Println("2 — Пополнить баланс")
		fmt.Println("3 — Снять средства")
		fmt.Println("0 — Выход")

		fmt.Print("Ваш выбор: ")
		_, err := fmt.Scan(&action)

		if err != nil {
			fmt.Println("Ошибка ввода")
			return
		}

		switch action {
		case 1:
			fmt.Printf("Текущий баланс: %.2f\n", balance)

		case 2:
			fmt.Print("Введите сумму пополнения: ")
			_, err := fmt.Scan(&amount)

			if err != nil || amount <= 0 {
				fmt.Println("Ошибка: некорректная сумма")
				return
			}

			balance += amount
			fmt.Printf("Баланс: %.2f\n", balance)

		case 3:
			fmt.Print("Введите сумму снятия: ")
			_, err := fmt.Scan(&amount)

			if err != nil || amount <= 0 {
				fmt.Println("Ошибка: некорректная сумма")
				return
			}

			if amount > balance {
				fmt.Println("Ошибка: недостаточно средств")
			} else {
				balance -= amount
				fmt.Printf("Баланс: %.2f\n", balance)
			}

		case 0:
			fmt.Println("До свидания!")
			return

		default:
			fmt.Println("Неизвестная операция")
		}
	}
}
