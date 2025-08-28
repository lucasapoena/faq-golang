package main

import "fmt"

func main() {
	var nome string = "Lucas"
	idade := 28
	const Empresa = "ApoenaTech"

	fmt.Printf("Nome: %s, Idade: %d, Empresa: %s\n", nome, idade, Empresa)

	// iota exemplo
	const (
		Janeiro = iota + 1
		Fevereiro
		Marco
	)
	fmt.Println("Janeiro =", Janeiro, "Fevereiro =", Fevereiro, "Março =", Marco)
}
