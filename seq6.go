package main

import "fmt"

func main() {
	var (
		salario float64
		nome    string
	)
	fmt.Println("Olá, qual seu nome?")
	fmt.Scan(&nome)
	fmt.Println("Senhor(a) ", nome, "informe seu salário para receber um aumento")
	fmt.Print("R$")
	fmt.Scan(&salario)
	aumento := salario + (salario * 15 / 100)
	fmt.Print(nome, " seu salário após o aumento será de: ", aumento)

}
