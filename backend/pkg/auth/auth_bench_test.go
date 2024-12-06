package auth

import (
	"testing"
	"time"
)

const (
	username       = "testuser"
	role           = "admin"
	expirationTime = 1 * time.Hour
)

var secretKey string = getSecretKeyFromEnv()

func BenchmarkGenerateJWT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		id:=1
		_, err := GenerateJWT(username, role,id, secretKey, expirationTime)
		if err != nil {
			b.Fatalf("failed to generate JWT: %v", err)
		}
	}
}

func BenchmarkValidateJWT(b *testing.B) {
	id:=1
	token, err := GenerateJWT(username, role, id,secretKey, expirationTime)
	if err != nil {
		b.Fatalf("failed to generate JWT for benchmarking: %v", err)
	}

	b.ResetTimer() // Сброс таймера перед началом бенчмарка
	for i := 0; i < b.N; i++ {
		_, err := ValidateJWT(token, secretKey)
		if err != nil {
			b.Fatalf("failed to validate JWT: %v", err)
		}
	}
}
