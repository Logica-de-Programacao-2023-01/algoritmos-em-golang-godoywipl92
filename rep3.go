package main

import "fmt"

func main() {
	var (
		i int
	)
	i = 1
	for i <= 19 {
		fmt.Println(i)
		i += 2
	}
}
