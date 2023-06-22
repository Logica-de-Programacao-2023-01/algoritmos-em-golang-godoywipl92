package main

import "fmt"

func main() {
	var (
		nome   string
		peso   float64
		altura float64
		sexo   string
	)
	fmt.Println("CALCULADORA DE IMC")
	fmt.Println("Olá, qual seu nome?")
	fmt.Scan(&nome)
	fmt.Println("Informe seu sexo:")
	fmt.Scan(&sexo)
	fmt.Println(nome, ", informe seu peso em quilos")
	fmt.Scan(&peso)
	fmt.Println("Informe sua altura em metros")
	fmt.Scan(&altura)
	imc := peso / (altura * altura)
	if imc < 18.5 {
		fmt.Print(nome, " você está abaixo do peso normal, com IMC de: ", imc)
	} else if 18.5 <= imc && imc <= 24.9 {
		fmt.Print(nome, " você está no peso ideal, com IMC de: ", imc)
	} else if 25 <= imc && imc <= 29.9 {
		fmt.Print(nome, " você está com sobrepeso, com IMC de: ", imc)
	} else if imc >= 30 {
		fmt.Print(nome, " você está com obesidade, com IMC de: ", imc, " CUIDADO!")
	}
}
