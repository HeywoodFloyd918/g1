package main

import (
	"fmt"
	"net/http"

	"github.com/HeywoodFloyd918/g1/liba"
	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("fmt")
	fmt.Println(liba.Sum(4, 1))
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}
