package main

import (
	
	"net/http"
	"os"
)

var secretKey = os.Getenv("secretKey")


func main() {
	http.HandleFunc("GET /",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hhhhhhhh"))
	})
	http.ListenAndServe(":8080", nil)
}
