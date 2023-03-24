package main

import "fmt"

func main() {
	var (
		numero1 int
		nome    string
	)
	fmt.Println("qual é o seu nome? ")
	fmt.Scan(&nome)
	fmt.Println("digite o numero do seu interesse para ser calculado o seu dobro , triplo , quadruplo :")
	fmt.Scan(&numero1)
	dobro := numero1 * 2
	triplo := numero1 * 3
	quadruplo := numero1 * 4
	fmt.Println("o dobro é ", dobro, "já o triplo é", triplo, "além disso o quadruplo é", quadruplo)
}
