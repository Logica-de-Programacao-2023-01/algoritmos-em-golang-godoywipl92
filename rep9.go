package main

import "fmt"

func main() {
	var (
		numero    int
		soma      int
		quantidade int
	)

	fmt.Println("Informe os números inteiros (digite 0 para encerrar):")

	for {
		fmt.Scan(&numero)

		if numero == 0 {
			break
		}

		soma += numero
		quantidade++
	}

	if quantidade > 0 {
		media := float64(soma) / float64(quantidade)
		fmt.Printf("Média: %.2f\n", media)
	} else {
		fmt.Println("Nenhum número foi informado.")
	}
}
