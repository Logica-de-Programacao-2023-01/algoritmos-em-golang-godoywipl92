package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Informe um número: ")
	fmt.Scan(&numero)

	fmt.Println("Tabuada de multiplicação para o número", numero)

	for i := 1; i <= 10; i++ {
		resultado := numero * i
		fmt.Printf("%d x %d = %d\n", numero, i, resultado)
	}
}
