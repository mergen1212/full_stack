package repository

import (
	"apiapp/internal/models"
	"apiapp/pkg/auth"

	"database/sql"
)

type UserRepo interface {
	CreateUser(user models.UserReg) (int, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(id int) (models.User, error)
	GetUserByName(name string) (models.User, error)
	UpdateUser(user models.User) error
	GetAllUsers() ([]models.User, error)
	DeleteUser(id int) error
}

type Repository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateUser(user models.UserReg) (int, error) {
	user.Pass = auth.Hash(user.Pass)
	query := `INSERT INTO users (username, email, password_hash) VALUES ($1,$2,$3) RETURNING id`
	var id int
	err := r.db.QueryRow(query, user.Name, user.Email, user.Pass).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password_hash FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.HashPass)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *Repository) GetUserByID(id int) (models.User, error) {
	var user models.User
	query := `SELECT id, username, email, hash_password FROM users WHERE id = ?`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.HashPass)
	if err != nil {
		return models.User{}, err

	}
	return user, nil

}

func (r *Repository) GetUserByName(name string) (models.User, error) {
	var user models.User
	query := `SELECT id, username, email, hash_password FROM users WHERE name = ?`
	err := r.db.QueryRow(query, name).Scan(&user.ID, &user.Name, &user.Email, &user.HashPass)
	if err != nil {
		return models.User{}, err
	}
	return user, nil

}
func (r *Repository) UpdateUser(user models.User) error {
	query := `UPDATE users SET username = ?,  email = ?, password_hash = ? WHERE id = ?`
	_, err := r.db.Exec(query, user.Name, user.Img, user.Email, user.HashPass, user.ID)
	if err != nil {
		return err
	}
	return nil

}

func (r *Repository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	query := `SELECT id, username,  email, password_hash FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.HashPass)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *Repository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
