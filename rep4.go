package main

import "fmt"

func main() {
	var (
		i int
	)
	i = 0
	for i <= 30 {
		fmt.Println(i)
		i += 3
	}
}
