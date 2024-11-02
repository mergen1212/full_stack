package main

import (
	"apiapp/internal/handlers"
	"net/http"
)

func main() {
	r:=http.NewServeMux()
	r.HandleFunc("GET /items", handlers.GetItems)

	err := http.ListenAndServe("0.0.0.0:8080", r)
	if err != nil {
		panic(err)
	}
}
