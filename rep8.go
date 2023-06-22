package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Informe um número inteiro positivo: ")
	fmt.Scan(&numero)

	fmt.Println("Divisores de", numero, ":")

	for i := 1; i <= numero; i++ {
		if numero%i == 0 {
			fmt.Println(i)
		}
	}
}
