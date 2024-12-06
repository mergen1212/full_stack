package repository

import (
	"apiapp/internal/models"
	"context"
	"database/sql"
	"testing"

	_ "github.com/lib/pq" // Импорт драйвера PostgreSQL
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestTokenRepository(t *testing.T) {
	ctx := context.Background()

	// Создание контейнера PostgreSQL
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpassword",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("Failed to start container: %s", err)
	}
	defer postgresC.Terminate(ctx)

	host, err := postgresC.Host(ctx)
	if err != nil {
		t.Fatalf("Failed to get host: %s", err)
	}
	port, err := postgresC.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatalf("Failed to get mapped port: %s", err)
	}

	dsn := "host=" + host + " port=" + port.Port() + " user=testuser password=testpassword dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err)
	}
	defer db.Close()

	// Создание таблицы токенов
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);


-- Create tokens table
CREATE TABLE IF NOT EXISTS tokens (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    token TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);`)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}
	repos:=NewUserRepo(db)
	repo := NewTokenRepo(db)

	// Тестирование CreateUser
	userReg := models.UserReg{Name: "testuser", Email: "test@example.com", Pass: "password"}
	id, err := repos.CreateUser(userReg)
	if err != nil {
		t.Fatalf("Failed to create user: %s", err)
	}

	// Тестирование GetUserByEmail
	user, err := repos.GetUserByEmail("test@example.com")
	if err != nil {
		t.Fatalf("Failed to get user by email: %s", err)
	}
	if user.ID != id {
		t.Errorf("Expected user ID %d, got %d", id, user.ID)
	}
	// Тестирование CreateToken
	token := Token{UserID: 1, Token: "testtoken"}
	err = repo.CreateToken(token)
	if err != nil {
		t.Fatalf("Failed to create token: %s", err)
	}

	// Тестирование GetTokenByValue
	retrievedToken, err := repo.GetTokenByValue("testtoken")
	if err != nil {
		t.Fatalf("Failed to get token by value: %s", err)
	}
	if retrievedToken.Token != token.Token {
		t.Errorf("Expected token %s, got %s", token.Token, retrievedToken.Token)
	}

	// Тестирование GetTokensByUserID
	tokens, err := repo.GetTokensByUserID(1)
	if err != nil {
		t.Fatalf("Failed to get tokens by user ID: %s", err)
	}
	if len(tokens) == 0 {
		t.Fatal("Expected to get tokens, got none")
	}

	// Тестирование DeleteToken
	err = repo.DeleteToken("testtoken")
	if err != nil {
		t.Fatalf("Failed to delete token: %s", err)
	}

	// Проверка, что токен был удален
	_, err = repo.GetTokenByValue("testtoken")
	if err == nil {
		t.Fatal("Expected error when getting deleted token, got none")
	}

	// Тестирование DeleteUserTokens
	err = repo.CreateToken(Token{UserID: 1, Token: "token1"})
	if err != nil {
		t.Fatalf("Failed to create token: %s", err)
	}
	err = repo.CreateToken(Token{UserID: 1, Token: "token2"})
	if err != nil {
		t.Fatalf("Failed to create token: %s", err)
	}

	err = repo.DeleteUserTokens(1)
	if err != nil {
		t.Fatalf("Failed to Delete token: %s", err)
	}
}
