package main

import (
	"fmt"
)

type Pessoa struct {
	nome string
	idade int
	salario int
}

func main(){
	pessoa := &Pessoa {
		nome: "Augusto",
		idade: 23,
		salario: 1000,
	}

	addSalary(pessoa, 10)
	fmt.Println(pessoa.salario)

}

func setName(name string) {
	fmt.Println("Hello", name)
}

func addSalary(p *Pessoa, bonus int){
	p.salario += bonus
}
