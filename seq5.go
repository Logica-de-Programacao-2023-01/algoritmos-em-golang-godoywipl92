package main

import "fmt"

func main() {
	var (
		nome   string
		idadea int
	)
	fmt.Println("Olá, qual seu nome?")
	fmt.Scan(&nome)
	fmt.Println("Sernhor(a) ", nome, "informe sua idade em anos e será informado sua idade em dias")
	fmt.Scan(&idadea)
	idaded := idadea * 365
	fmt.Print("Sua idade em dias é: ", idaded)
}
