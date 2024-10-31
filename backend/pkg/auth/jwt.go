package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims представляет собой структуру для хранения пользовательских данных в токене.
type JWTClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateJWT создает новый JWT токен для указанного пользователя.
func GenerateJWT(username string,role string, secretKey string, expirationTime time.Duration) (string, error) {

	claims := &JWTClaims{
		Username: username,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ValidateJWT проверяет действительность JWT токена.
func ValidateJWT(tokenString string, secretKey string) (*JWTClaims, error) {
	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Проверка метода подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
