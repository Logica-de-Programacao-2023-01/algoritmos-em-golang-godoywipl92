package main

import "fmt"

func main() {
	var (
		num1 int
		num2 int
		soma int
		mult int
	)

	fmt.Println("Infome 2 números")
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	if num1 >= 0 && num2 >= 0 {
		mult = num1 * num2
		fmt.Println("A múltiplicação é:", mult)
	} else {
		soma = num1 + num2
		fmt.Println("A soma é:", soma)
	}
}
