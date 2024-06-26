// Server2 é um servidor mínimo de "eco" e contador
/* Inicie com o comando go run server2.go &
   go run fetch.go http://localhost:8000/help
   go run fetch.go http://localhost:8000/count
   ps -aux | grep server2 para ver o processo
   killall -9 server2 para matar o processo
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler ecoa o componente Path do URL requisitado
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter ecoa o número de chamadas até agora
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
