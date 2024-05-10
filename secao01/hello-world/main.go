package main

import (
	"fmt"
)

func main(){
	name, salario := "Augusto", 1000
	setName(name)
	newSalary, bonus := addSalary(salario, 10)
	fmt.Println("Novo Salário:", newSalary)
	fmt.Println("Bônus:", bonus)
}

func setName(name string) {
	fmt.Println("Hello", name)
}

func addSalary(valorAtual int, bonus int) (int, int){
	return valorAtual + bonus, bonus // Dois retornos
}
