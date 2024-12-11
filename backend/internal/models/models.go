package models

import (
    "time"
)
//
type User struct {
    ID          int       `json:"id"`
    Username    string    `json:"username"`
    Email       string    `json:"email"`
    PasswordHash string    `json:"pass"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type Token struct {
    ID        int       `json:"id"`
    UserID    int       `json:"user_id"`
    Token     string    `json:"token"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
