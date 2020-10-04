package main

import "fmt"

func main() {
	var val = somaDoisCincos()
	fmt.Printf("A soma de 5+5 é: %d\n", val)
}

func somaDoisCincos() int {
	var soma = 5 + 5
	return soma
}
