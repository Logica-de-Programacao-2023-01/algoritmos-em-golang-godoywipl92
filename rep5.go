package main

import "fmt"

func main() {
	var (
		i int
	)
	i = 10
	for i >= 0 {
		fmt.Println(i)
		i--
	}
}
