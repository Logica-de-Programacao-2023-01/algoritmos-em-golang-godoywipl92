package main

import "fmt"

func main() {
	var (
		numero1 int
		numero2 int
		numero3 int
	)
	fmt.Println("digite o primeiro numero :")
	fmt.Scan("&numero1")
	fmt.Println("digite o segundo numero :")
	fmt.Scan("numero2")
	fmt.Println("digite o terceiro numero :")
	fmt.Scan("numero3")
	soma := numero1 + numero2 + numero3
	fmt.Println("a soma é ", soma)
}
