package main

import (
	"fmt"
)

func main(){
	salFuncs := make(map[string] int)
	salFuncs["Augusto"] = 10
	salFuncs["Fadanelli"] = 20

	sal, exists := salFuncs["Fadanelli"]
	fmt.Println(sal, exists)
	qtdSal := len(salFuncs)
	fmt.Println(qtdSal)
}

