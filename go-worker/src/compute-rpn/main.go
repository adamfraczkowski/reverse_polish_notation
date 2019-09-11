package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func computeRPN(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/compute", computeRPN).Methods("POST")

	http.ListenAndServe(":8080", r)
}
