package main

import (
	"apiapp/internal/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)


var jwtSecret = []byte("your_secret_key") // Замените на свой секретный ключ

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Хеширование пароля
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "could not hash password", http.StatusInternalServerError)
        return
    }
    user.PasswordHash = string(hashedPassword)

    // Сохранение пользователя в БД
    _, err = db.Exec("INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)", user.Username, user.Email, user.PasswordHash)
    if err != nil {
        http.Error(w, "could not create user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var user models.User
    var userpass string
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    userpass=user.PasswordHash
    // Поиск пользователя в БД
    row := db.QueryRow("SELECT id, password_hash FROM users WHERE username = $1", user.Username)
    err := row.Scan(&user.ID, &user.PasswordHash)
    if err != nil {
        if err == sql.ErrNoRows {
        	log.Print(1)
            http.Error(w, "invalid username or password", http.StatusUnauthorized)
            return
        }
        http.Error(w, "could not query user", http.StatusInternalServerError)
        return
    }

    // Проверка пароля
    err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(userpass))
    if err != nil {
    	log.Print(2)
        http.Error(w, "invalid username or password", http.StatusUnauthorized)
        return
    }

    // Генерация JWT
    token, err := generateJWT(user.ID)
    if err != nil {
        http.Error(w, "could not generate token", http.StatusInternalServerError)
        return
    }

    // Сохранение токена в БД
    _, err = db.Exec("INSERT INTO tokens (user_id, token) VALUES ($1, $2)", user.ID, token)
    if err != nil {
        http.Error(w, "could not save token", http.StatusInternalServerError)
        return
    }

    // Возвращаем токен пользователю
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func generateJWT(userID int) (string, error) {
    claims := &jwt.MapClaims{
        "sub": userID,
        "exp": time.Now().Add(time.Hour * 72).Unix(), // Токен будет действителен 72 часа
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}


var db *sql.DB

func initDB() {
    var err error
    connStr := "postgresql://user:password@postgres:5432/dbname?sslmode=disable" // замените на свои данные
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    // Создание таблиц
    createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        password_hash VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT NOW(),
        updated_at TIMESTAMP DEFAULT NOW()
    );`

    createTokensTable := `
    CREATE TABLE IF NOT EXISTS tokens (
        id SERIAL PRIMARY KEY,
        user_id INT NOT NULL,
        token TEXT UNIQUE NOT NULL,
        created_at TIMESTAMP DEFAULT NOW(),
        updated_at TIMESTAMP DEFAULT NOW(),
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );`

    _, err = db.Exec(createUsersTable)
    if err != nil {
        log.Fatal(err)
    }

    _, err = db.Exec(createTokensTable)
    if err != nil {
        log.Fatal(err)
    }
}




func main() {
    initDB() // Инициализация базы данных

    http.HandleFunc("/register", Register)
    http.HandleFunc("/login", Login)

    log.Println("Сервер запущен на порту :8080")
    log.Fatal(http.ListenAndServe(":8080", nil)) // Запуск сервера на порту 8080
}
