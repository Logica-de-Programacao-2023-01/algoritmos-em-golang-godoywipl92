package main

import "fmt"

func main() {
	var (
		salario      float64
		novo_salario float64
	)
	fmt.Println("Informe o salário")
	fmt.Scan(&salario)
	if salario < 1000 {
		novo_salario = salario + (salario * 10 / 100)
	} else {
		novo_salario = salario + (salario * 5 / 100)
	}
	fmt.Printf("O novo salário é de R$ %.2f", novo_salario)
}
