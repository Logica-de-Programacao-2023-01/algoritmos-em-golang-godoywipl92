package main

import "fmt"

func main() {
	var (
		nome   string
		dias   float64
		diaria float64
	)
	fmt.Println("Olá, qual seu nome?")
	fmt.Scan(&nome)
	fmt.Println(nome, " quantos dias por mês o(a) senhor(a) trabalha?")
	fmt.Scan(&dias)
	fmt.Println("Qual valor pago por dia?")
	fmt.Print("R$")
	fmt.Scan(&diaria)
	salario := dias * diaria
	fmt.Print("O valor do seu salário é de: R$", salario)
}
