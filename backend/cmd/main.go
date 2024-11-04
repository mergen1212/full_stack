package main

import (
	"apiapp/internal/handlers"
	"apiapp/internal/repository"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var secretKey = os.Getenv("secretKey")


func main() {
	url := os.Getenv("TOKEN")
	if url==""{
		panic("uykuykyyukukkukujhtg")
	}
	log.Print(url)
	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	defer db.Close()
	
	
	repo := repository.NewUserRepo(db)
	h := handlers.NewHandler("ccc", repo)
	r := h.InitRoutes()
	http.ListenAndServe(":8080", r)
}
