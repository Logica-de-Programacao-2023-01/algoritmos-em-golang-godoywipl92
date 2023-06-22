package main

import "fmt"

func main() {
	var pesoKg float64

	fmt.Print("Informe o peso em quilos: ")
	fmt.Scan(&pesoKg)

	pesoLb := pesoKg * 2.20462

	fmt.Println("O peso em libras é:", pesoLb)
}
