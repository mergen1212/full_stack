package router

import (
	
	"apiapp/internal/service"
	"net/http"
)

func NewRouterUser(serv *service.Service) *http.ServeMux {
	r:=http.NewServeMux()
	r.HandleFunc("/live",serv.Head)
	r.HandleFunc("POST /create",serv.Create)
	r.HandleFunc("GET /readAll",serv.GetAll)
 	r.HandleFunc("PUT /update",serv.Update)
 	r.HandleFunc("DELETE /delete",serv.Delete)
	
	return r
}

func NewRouterToken(r *http.ServeMux,serv *service.TokenHandler){
	r.HandleFunc("POST /token",serv.CreateToken)

}
