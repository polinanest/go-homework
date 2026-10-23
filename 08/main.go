package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i <= n/i; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var n int

	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&n)

	if err != nil {
		fmt.Println("Ошибка: введите целое число")
		return
	}

	if isPrime(n) {
		fmt.Println("Число простое")
	} else {
		fmt.Println("Число не является простым")
	}
}
