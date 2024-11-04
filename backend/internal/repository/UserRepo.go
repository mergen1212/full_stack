package repository

import (
	"apiapp/internal/models"
	"log"

	"database/sql"
)

type UserRepo interface {
	CreateUser(user models.UserReg) error
	GetUserByEmail(email string) (models.User, error)
	GetUserByID(id int) (models.User, error)
	GetUserByName(name string)(models.User, error)
	UpdateUser(user models.User)(error)
}

type Repository struct {
	db *sql.DB
}
func NewUserRepo(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateUser(user models.UserReg) error {
	var newUser models.User
	query := `INSERT INTO users (name,img, email, hash_password) VALUES (?,?,?,?)`
	_,err := r.db.Exec(query, user.Name,"", user.Email, user.Pass)
	log.Print("user",newUser)
	if err != nil {
		return err
	}
	return nil
}



func (r *Repository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := `SELECT id, name, email, hash_password FROM users WHERE email = ?`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.HashPass)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *Repository) GetUserByID(id int) (models.User, error) {
	var user models.User
	query := `SELECT id, name, email, hash_password FROM users WHERE id = ?`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.HashPass)
	if err != nil {
		return models.User{}, err

	}
	return user, nil

}

func (r *Repository) GetUserByName(name string)(models.User, error){
	var user models.User
 	query := `SELECT id, name, email, hash_password FROM users WHERE name = ?`
 	err := r.db.QueryRow(query, name).Scan(&user.ID, &user.Name, &user.Email, &user.HashPass)
 	if err != nil {
 		return models.User{}, err
 	}
 	return user, nil
 
}
func (r *Repository) UpdateUser(user models.User)(error){
	query := `UPDATE users SET name = ?, img = ?, email = ?, hash_password = ? WHERE id = ?`
 	_, err := r.db.Exec(query, user.Name, user.Img, user.Email, user.HashPass, user.ID)
 	if err != nil {
 		return err
 	}
 	return nil
 
}