package main

import (
	"fmt"
)

func main(){
	/* ////////////
	   // Arrays //
	*/ ////////////

	// Alternativa 1
	var salarios1 [5]int // Array de inteiros com cinco posições
	salarios1[0] = 1000
	salarios1[1] = 1500
	// range retorna o indice e o valor.
	// _ é usado quando não queremos guardar o valor retornado.
	for _, salario := range salarios1{
		fmt.Println(salario)
	}
	fmt.Println()

	// Alternativa 2
	salarios2 := make([]int, 5) // Array de inteiros com cinco posições
	salarios2[0] = 2000
	salarios2[1] = 3000
	salarios2[4] = 6000
	for _, salario := range salarios2{
		fmt.Println(salario)
	}
	fmt.Println()

	/* ////////////
	   // Slices //
	*/ ////////////
	
	// Valores pré-definidos
	salarios3 := []int{1000, 2000, 3000, 4000, 5000}
	for i := 0; i < len(salarios3); i++{
		salarios3[i] += 100 + i
	}
	for _, salario := range salarios3{
		fmt.Println(salario)
	}
	fmt.Println()

	// Slice vazio
	salarios4 := []int{}
	for i := 0; i < 5; i++{
		salarios4 = append(salarios4, 100 + i)
	}
	for _, salario := range salarios4{
		fmt.Println(salario)
	}

}

