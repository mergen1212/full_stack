package repository

import (
	"database/sql"
	"time"
)

type Token struct {
	ID        int
	UserID    int
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TokenRepo interface {
	CreateToken(token Token) error
	GetTokenByValue(tokenValue string) (Token, error)
	GetTokensByUserID(userID int) ([]Token, error)
	DeleteToken(tokenValue string) error
	DeleteUserTokens(userID int) error
}

type TokenRepository struct {
	db *sql.DB
}

func NewTokenRepo(db *sql.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

func (r *TokenRepository) CreateToken(token Token) error {
	query := `INSERT INTO tokens (user_id, token) VALUES ($1, $2)`
	_, err := r.db.Exec(query, token.UserID, token.Token)
	return err
}

func (r *TokenRepository) GetTokenByValue(tokenValue string) (Token, error) {
	var token Token
	query := `SELECT id, user_id, token, created_at, updated_at FROM tokens WHERE token = $1`
	err := r.db.QueryRow(query, tokenValue).Scan(&token.ID, &token.UserID, &token.Token, &token.CreatedAt, &token.UpdatedAt)
	return token, err
}

func (r *TokenRepository) GetTokensByUserID(userID int) ([]Token, error) {
	var tokens []Token
	query := `SELECT id, user_id, token, created_at, updated_at FROM tokens WHERE user_id = $1`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var token Token
		err := rows.Scan(&token.ID, &token.UserID, &token.Token, &token.CreatedAt, &token.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (r *TokenRepository) DeleteToken(tokenValue string) error {
	query := `DELETE FROM tokens WHERE token = $1`
	_, err := r.db.Exec(query, tokenValue)
	return err
}

func (r *TokenRepository) DeleteUserTokens(userID int) error {
	query := `DELETE FROM tokens WHERE user_id = $1`
	_, err := r.db.Exec(query, userID)
	return err
}