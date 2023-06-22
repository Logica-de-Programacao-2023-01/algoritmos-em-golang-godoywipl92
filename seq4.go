package main

import "fmt"

func main() {
	var (
		nome string
		num1 float64
		num2 float64
		num3 float64
	)
	fmt.Println("Olá, qual seu nome?")
	fmt.Scan(&nome)
	fmt.Println("Sr(a). ", nome, " Informe 3 números para informarmos a sua média ponderada")
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	fmt.Scan(&num3)
	media := (num1*2 + num2*3 + num3*5) / 10
	fmt.Print("A média dos números ", num1, ", ", num2, " e ", num3, " é ", media)
}
