package main

import "fmt"

func main() {
	fmt.Println(Add(1, 2))
}

func Add(a, b int) int {
	var c = a
	var d = b

	return c + d
}

func unused() {
	
}
