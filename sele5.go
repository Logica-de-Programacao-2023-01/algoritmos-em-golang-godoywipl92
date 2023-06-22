package main

import "fmt"

func main() {
	var (
		num   int
		mult3 int
		mult5 int
	)

	fmt.Println("Informe um número")
	fmt.Scan(&num)
	mult3 = num % 3
	mult5 = num % 5

	if mult3 == 0 {
		if mult5 == 0 {
			fmt.Printf("%d é múltiplo de 3 e 5 simultaneamente", num)
		} else {
			fmt.Printf("%d é múltiplo de 3 mas não é múltiplo de 5", num)
		}
	} else {
		fmt.Printf("%d não é múliplo", num)
	}
}
