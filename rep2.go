package main

import "fmt"

func main() {
	var (
		i int
	)
	i = 0
	for i <= 20 {
		fmt.Println(i)
		i += 2
	}
}
