package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	// Verifica quantidade de argumentos
	if len(os.Args) != 2 {
		fmt.Println("Erro, quantidade inválida de argumentos!")
		os.Exit(1)
	}

	// Verifica se o argumento é um número inteiro
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Erro, não é um número válido!")
		os.Exit(1)
	}
	fmt.Println("Número convertido:", n)

	// Verifica se o arquivo existe
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
}

