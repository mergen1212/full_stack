package main

import (
	"apiapp/internal/handlers"
	"net/http"
)

func main() {
	r:=http.NewServeMux()
	r.HandleFunc("/ws", handlers.HandlerFuncWS)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
