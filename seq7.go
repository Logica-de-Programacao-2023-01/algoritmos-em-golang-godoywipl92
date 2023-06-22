package main

import "fmt"

func main() {

	var (
		nome string
		num  int
	)
	fmt.Println("Olá, qual seu nome?")
	fmt.Scan(&nome)
	fmt.Println(nome, ", informe um número inteiro")
	fmt.Scan(&num)
	sucessor := num + 1
	antecessor := num - 1
	fmt.Print("O sucessor de ", num, " é: ", sucessor, ", e seu antecessor é: ", antecessor)
}
