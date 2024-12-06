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

func TestUserRepository(t *testing.T) {
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

	// Создание таблицы пользователей
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
    expires_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);`)
	if err != nil {
		t.Fatalf("Failed to create table: %s", err)
	}

	repo := NewUserRepo(db)
	
	// Тестирование CreateUser
	userReg := models.UserReg{Name: "testuser", Email: "test@example.com", Pass: "password"}
	id, err := repo.CreateUser(userReg)
	if err != nil {
		t.Fatalf("Failed to create user: %s", err)
	}

	// Тестирование GetUserByEmail
	user, err := repo.GetUserByEmail("test@example.com")
	if err != nil {
		t.Fatalf("Failed to get user by email: %s", err)
	}
	if user.ID != id {
		t.Errorf("Expected user ID %d, got %d", id, user.ID)
	}

	// Тестирование DeleteUser
	err = repo.DeleteUser(id)
	if err != nil {
		t.Fatalf("Failed to delete user: %s", err)
	}

	// Проверка, что пользователь был удален
	_, err = repo.GetUserByEmail("test@example.com")
	if err == nil {
		t.Fatal("Expected error when getting deleted user, got none")
	}
}
