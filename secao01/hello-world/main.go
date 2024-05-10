package main

import (
	"fmt"
)

func main(){
	name := getName()
	idade := 23
	fmt.Println("Hello", name, idade)
}

func getName() string{
	return "Augusto"
}
