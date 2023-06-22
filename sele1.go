package main

import "fmt"

func main() {
	var (
		num1 int
		num2 int
	)
	fmt.Println("Olá usuário, favor informe 2 números")
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	if num1 > num2 {
		fmt.Println(num1, " é maior que ", num2)
		fmt.Println("R= ", num1)
	} else if num2 > num1 {
		fmt.Println(num2, " é maior que ", num1)
		fmt.Println("R= ", num2)
	} else if num1 == num2 {
		fmt.Println(num1, " é igual a ", num2)
		fmt.Println("R= ", num1, "=", num2)
	}
}
