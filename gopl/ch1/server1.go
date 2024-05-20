// Server1 é um servidor de "eco" mínimo.
/* Inicie com o comando go run server1.go &
   go run fetch.go http://localhost:8000
   go run fetch.go http://localhost:8000/help
   ps -aux | grep server1 para ver o processo
   killall -9 server1 para matar o processo
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // Cada requisição chama handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Handler ecoa o componente Path do URL requisistado
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
