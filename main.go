package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("GET /", handleGet)
	http.ListenAndServe(":8080", nil)

}

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Funcionou!"))
}
