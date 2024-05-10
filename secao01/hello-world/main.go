package main

import (
	"fmt"
	"os"
)

func main(){
	// Args é um vetor de duas posições: Args[0,1]
	// Args[0]: Slice, nome do próprio arquivo
	// Args[1]: Valor informado
	if len(os.Args) !=2 { 
		os.Exit(1) // Se for passado mais de dois argumentos, o programa será fechado.
	}

	fmt.Println("Hello", os.Args[1])
}
