package main

import "fmt"

func main() {
	var (
		num1 int
		num2 int
		num3 int
	)
	//estrutural
	fmt.Println("Olá usuário, está interface será usada para evidenciar o menor número de 3 escolhidos pelo senhor")
	fmt.Println("Por favor, informe o primeiro número")
	fmt.Scan(&num1)
	fmt.Println("Ok, agora o segundo")
	fmt.Scan(&num2)
	fmt.Println("Perfeito. Por último, o terceiro número")
	fmt.Scan(&num3)
	//parcelas if
	if num1 < num2 && num1 < num3 {
		fmt.Println(num1, " é o menor número")
	} else if num2 < num1 && num2 < num3 {
		fmt.Println(num2, " é o menor número")
	} else if num3 < num1 && num3 < num2 {
		fmt.Println(num3, " é o menor número")
	} else {
		fmt.Println("Os 3 núemeros são iguais entre sí")
	}
}
