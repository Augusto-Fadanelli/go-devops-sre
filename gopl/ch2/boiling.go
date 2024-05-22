// Boiling exibe o ponto de ebulição da água.
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Ponto de ebulição = %g °F or %g °C\n", f, c)
}
