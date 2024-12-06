package service

import (
    "encoding/json"
    "net/http"
    "strconv"
)

type TokenHandler struct {
    service *TokenService
}

func NewTokenHandler(service *TokenService) *TokenHandler {
    return &TokenHandler{service: service}
}

type TokenResponse struct {
    Token string `json:"token"`
}



func (h *TokenHandler) CreateToken(w http.ResponseWriter, r *http.Request) {
    var userID int
    if err := json.NewDecoder(r.Body).Decode(&userID); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    token, err := h.service.CreateToken(userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    response := TokenResponse{Token: token}
    json.NewEncoder(w).Encode(response)
    w.WriteHeader(http.StatusCreated)
}

func (h *TokenHandler) ValidateToken(w http.ResponseWriter, r *http.Request) {
    tokenValue := r.URL.Query().Get("token")
    token, err := h.service.ValidateToken(tokenValue)
    if err != nil {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }
	
    json.NewEncoder(w).Encode(token)
}

func (h *TokenHandler) GetUserTokens(w http.ResponseWriter, r *http.Request) {
    userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    tokens, err := h.service.GetUserTokens(userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(tokens)
}

func (h *TokenHandler) RevokeToken(w http.ResponseWriter, r *http.Request) {
    tokenValue := r.URL.Query().Get("token")
    err := h.service.RevokeToken(tokenValue)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func (h *TokenHandler) RevokeAllUserTokens(w http.ResponseWriter, r *http.Request) {
    userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    err = h.service.RevokeAllUserTokens(userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
