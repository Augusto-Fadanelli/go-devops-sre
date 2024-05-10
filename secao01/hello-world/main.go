package main

import (
	"fmt"
)

type Pessoa struct {
	nome string
	idade int
	salario int
}
func (p *Pessoa) addSalary(bonus int){
	p.salario += bonus
}

func main(){
	pessoa := new(Pessoa)
	pessoa.nome = "Augusto"
	pessoa.idade = 23
	pessoa.salario = 1000

	pessoa.addSalary(10)
	fmt.Println(pessoa.salario)

}

func setName(name string) {
	fmt.Println("Hello", name)
}
