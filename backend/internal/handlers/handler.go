package handlers

import (
	"apiapp/internal/models"
	"apiapp/internal/repository"
	"apiapp/pkg/auth"
	"log"

	"net/http"

	"time"

	"github.com/goccy/go-json"
)

type Handler struct {
	secretKey string
	repo      repository.UserRepo
}

func NewHandler(secretKey string, repo repository.UserRepo) *Handler {
	return &Handler{
		secretKey: secretKey,
		repo:      repo,
	}
}
func (h *Handler) InitRoutes() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("GET /",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("jjjj"))
	})
	r.HandleFunc("POST /auth", h.AuthUser)
	r.HandleFunc("POST /reg", h.RegUser)
	r.HandleFunc("POST /update",h.UpdateUser)
	return r
}

func (h *Handler) AuthUser(w http.ResponseWriter, r *http.Request) {
	var User models.UserAuth
	json.NewDecoder(r.Body).Decode(&User)
	h.repo.GetUserByEmail(User.Email)
	t, err := auth.GenerateJWT(User.Email, "user", h.secretKey, 10*time.Minute)
	if err != nil {
		http.Error(w, "Ошибка при генерации токена", http.StatusInternalServerError)
		return
	}
	auth.AddJWTCookie(w, t)
	
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) RegUser(w http.ResponseWriter, r *http.Request) {
	var User models.UserReg
	var err error
	json.NewDecoder(r.Body).Decode(&User)
	User.Pass=auth.Hash(User.Pass)
	err=h.repo.CreateUser(User)
	if err!=nil{
		log.Print(err)
		http.Error(w, "Ошибка при CreateUser", http.StatusInternalServerError)
		return
	}
	t, err := auth.GenerateJWT(User.Email, "user", h.secretKey, 10*time.Minute)
	if err != nil {
		http.Error(w, "Ошибка при генерации токена", http.StatusInternalServerError)
		return
	}
	auth.AddJWTCookie(w, t)
	w.WriteHeader(http.StatusOK)
}

func(h *Handler) UpdateUser (w http.ResponseWriter, r *http.Request){
	err,token:=auth.GetJWTCookie(r)
	if err!=nil{
		http.Error(w, "Ошибка взятия токен", http.StatusForbidden)
		return
	}
	jwtval,err:=auth.ValidateJWT(token,h.secretKey)
	if err!=nil{
		http.Error(w, "Ошибка не валидный токен", http.StatusForbidden)
		return
	}
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	if jwtval.Username!=user.Name{
		http.Error(w, "Ошибка не валидный токен", http.StatusForbidden)
		return
	}
	_,err=h.repo.GetUserByID(user.ID)
	if err!=nil{
		http.Error(w, "Ошибка не валидный токен", http.StatusForbidden)
		return
	}
	err=h.repo.UpdateUser(user)
	if err!=nil{
		http.Error(w, "Ошибка обновления", http.StatusForbidden)
		return
	}
}