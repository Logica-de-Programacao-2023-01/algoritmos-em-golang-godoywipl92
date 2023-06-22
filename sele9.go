package main

import "fmt"

func main() {
	var num1, num2, num3 float64

	fmt.Print("Informe o primeiro número: ")
	fmt.Scan(&num1)

	fmt.Print("Informe o segundo número: ")
	fmt.Scan(&num2)

	fmt.Print("Informe o terceiro número: ")
	fmt.Scan(&num3)

	// Ordenação dos números
	if num1 > num2 {
		num1, num2 = num2, num1
	}

	if num2 > num3 {
		num2, num3 = num3, num2

		if num1 > num2 {
			num1, num2 = num2, num1
		}
	}

	fmt.Println("Números em ordem crescente:", num1, num2, num3)
}
