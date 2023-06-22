package main

import "fmt"

func main() {
	var (
		preco float64
		nome  string
	)
	fmt.Println("Olá, qual seu nome?")
	fmt.Scan(&nome)
	fmt.Println("Informe o preço do produto para validarmos o desconto Sr(a). ", nome)
	fmt.Scan(&preco)
	desconto := preco - (preco * 10 / 100)
	fmt.Print("O valor do seu produto com desconto é de: R$ ", desconto)
}
