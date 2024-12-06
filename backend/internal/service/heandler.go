package service

import (
	"apiapp/internal/models"
	"apiapp/internal/repository"
	"apiapp/pkg/auth"
	"time"

	"encoding/json"
	"net/http"
	"strconv"
)

type Service struct {
	repository repository.UserRepo
	token repository.TokenRepo
}

func NewService(repository repository.UserRepo,token repository.TokenRepo) *Service {
	return &Service{
		repository: repository,
		token: token,
	}
}

func  (h *Service) Head(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from API!"))
}

func (h *Service) Create(w http.ResponseWriter, r *http.Request) {
	var item models.UserReg
	var id int
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id,err := h.repository.CreateUser(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token,err:=auth.GenerateJWT(item.Name,"user",id,"sic",time.Hour)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t:=repository.Token{
		UserID: id,
		Token: token,
	}
	h.token.CreateToken(t)
	auth.AddJWTCookie(w,token)
	w.WriteHeader(http.StatusCreated)
}

func (h *Service) GetAll(w http.ResponseWriter, r *http.Request) {
	items, err := h.repository.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}

func (h *Service) GetById(w http.ResponseWriter, r *http.Request) {
	id,err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item, err := h.repository.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func (h *Service) Update(w http.ResponseWriter, r *http.Request) {
	id,err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var item models.User
	if err = json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if item,err=h.repository.GetUserByID(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err := h.repository.UpdateUser(item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Service) Delete(w http.ResponseWriter, r *http.Request) {
	id := 0 // Parse id from request

	if err := h.repository.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
