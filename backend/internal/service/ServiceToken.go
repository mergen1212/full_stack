package service

import (
	"apiapp/internal/repository"
	"apiapp/pkg/auth"
	
	"time"
)

type TokenService struct {
	repo *repository.TokenRepository
	secret string
}

func NewTokenService(repo *repository.TokenRepository,sec string) *TokenService {
	return &TokenService{repo: repo,secret: sec}
}

func (s *TokenService) CreateToken(userID int) (string, error) {
	tokenValue, err := auth.GenerateJWT("user", "sic", userID, "user", time.Hour)
	if err != nil {
		return "", err
	}
	token := repository.Token{
		UserID:    userID,
		Token:     tokenValue,
	}
	return tokenValue, s.repo.CreateToken(token)
}


func (s *TokenService) ValidateToken(tokenValue string) (*auth.JWTClaims, error) {
	var err error
	var col *auth.JWTClaims
	col,err=auth.ValidateJWT(tokenValue,s.secret)
	if err!=nil{
		return nil,err
	}
	_,err=s.repo.GetTokenByValue(tokenValue)
	if err!=nil{
		return nil,err
	}
	return col,nil
}

func (s *TokenService) GetUserTokens(userID int) ([]repository.Token, error) {
	return s.repo.GetTokensByUserID(userID)
}

func (s *TokenService) RevokeToken(tokenValue string) error {
	return s.repo.DeleteToken(tokenValue)
}

func (s *TokenService) RevokeAllUserTokens(userID int) error {
	return s.repo.DeleteUserTokens(userID)
}
