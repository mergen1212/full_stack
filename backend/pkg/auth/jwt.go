package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims представляет собой структуру для хранения пользовательских данных в токене.
type JWTClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	ID       int    `json:"id"`
	jwt.RegisteredClaims
}

// GenerateJWT создает новый JWT токен для указанного пользователя.
func GenerateJWT(username string, role string,id int, secretKey string, expirationTime time.Duration) (string, error) {

	claims := &JWTClaims{
		Username: username,
		Role:     role,
		ID:       id,
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

func AddJWTCookie(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)
}
func GetJWTCookie(r *http.Request) (string,error){
	cookie,err:=r.Cookie("jwt")
	if err!=nil{
		return "",err
	}
	val:=cookie.Value
	return val,err
}