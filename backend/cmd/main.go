package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var secretKey = os.Getenv("secretKey")


func main() {
	http.HandleFunc("GET /",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w,"oio")
	})
	http.ListenAndServe(":8080", nil)
}
