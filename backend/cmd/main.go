package main

import (
	"apiapp/internal/repository"
	"apiapp/internal/router"
	"apiapp/internal/service"
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:postgres@postgres/fullapp?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err!=nil{
		panic(err)
	}
	repo:=repository.NewUserRepo(db)
	tok:=repository.NewTokenRepo(db)
	userserv:=service.NewService(repo,tok)

	r:=router.NewRouterUser(userserv)
	http.ListenAndServe(":8080", r)
}
