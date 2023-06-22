package main

import "fmt"

func main() {
	var (
		numero int
		maior  int
	)

	fmt.Println("Informe os números inteiros (digite 0 para encerrar):")

	for {
		fmt.Scan(&numero)

		if numero == 0 {
			break
		}

		if numero > maior {
			maior = numero
		}
	}

	if maior != 0 {
		fmt.Println("O maior número é:", maior)
	} else {
		fmt.Println("Nenhum número foi informado.")
	}
}
