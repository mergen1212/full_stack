package auth

import (
	"os"
	"testing"
	"time"
)

// получение secretKey из ENV
func getSecretKeyFromEnv() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey = "default_secret_key"
	}
	return secretKey
}

func TestGenerateJWT(t *testing.T) {
	secretKey := getSecretKeyFromEnv()
	username := "testuser"
	role := "admin"
	expirationTime := time.Hour

	token, err := GenerateJWT(username, role, secretKey, expirationTime)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if token == "" {
		t.Fatal("Expected a token, got an empty string")
	}
}

func TestValidateJWT(t *testing.T) {
	secretKey := getSecretKeyFromEnv()
	username := "testuser"
	role := "admin"
	expirationTime := time.Hour

	// Генерируем токен
	token, err := GenerateJWT(username, role, secretKey, expirationTime)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Проверяем токен
	claims, err := ValidateJWT(token, secretKey)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if claims.Username != username {
		t.Errorf("Expected username %s, got %s", username, claims.Username)
	}

	if claims.Role != role {
		t.Errorf("Expected role %s, got %s", role, claims.Role)
	}
}

func TestValidateJWTInvalidToken(t *testing.T) {
	secretKey := getSecretKeyFromEnv()
	invalidToken := "invalid.token.string"

	claims, err := ValidateJWT(invalidToken, secretKey)
	if err == nil {
		t.Fatal("Expected an error for invalid token, got none")
	}

	if claims != nil {
		t.Fatal("Expected claims to be nil for invalid token")
	}
}

func TestValidateJWTExpiredToken(t *testing.T) {
	secretKey := getSecretKeyFromEnv()
	username := "testuser"
	role := "admin"
	expirationTime := -time.Hour // Токен будет просрочен

	// Генерируем токен
	token, err := GenerateJWT(username, role, secretKey, expirationTime)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Проверяем токен
	claims, err := ValidateJWT(token, secretKey)
	if err == nil {
		t.Fatal("Expected an error for expired token, got none")
	}

	if claims != nil {
		t.Fatal("Expected claims to be nil for expired token")
	}
}
