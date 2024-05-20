package main

import (
	"fmt"
)

type Empregado struct {
	nome string
	id int
}

func main(){
	var inteiro = 45
	var ponteiro *int = &inteiro

	// Funcionamento básico de um ponteiro
	fmt.Println("Valor da variável inteiro:", inteiro)
	fmt.Println("Local da variável inteiro:", &inteiro)
	fmt.Println("Valor da variável ponteiro:", ponteiro)
	fmt.Println("Local da variável ponteiro:", &ponteiro)
	fmt.Println("Valor referenciado pela variável ponteiro:", *ponteiro)

	// Utilizando ponteiros em funções
	pointerFunction(&inteiro)
	fmt.Println("\nNovo valor da variável inteiro: ", inteiro)

	// Utilizando ponteiros em structs
	emp := Empregado{"Augusto", 123}
	pts := &emp

	fmt.Println("\nemp:", emp)
	fmt.Println("pts:", pts)
	
	pts.nome = "Fadanelli"
	fmt.Println(emp)

	//teste
	fmt.Println(pts)
	fmt.Println(*pts)
}

func pointerFunction(a *int) {
	*a = 400
}

