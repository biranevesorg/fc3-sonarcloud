package main

import "fmt"

func main() {
	fmt.Println(soma(2, 2))
}

func soma(a int, b int) int {
	return a + b
}

func subtracao(a int, b int) int {
	return a - b
}

func multiplicacao(a int, b int) int {
	return a * b
}

func divisao(a int, b int) int {
	return a / b
}
