package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func my_handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", my_handler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
