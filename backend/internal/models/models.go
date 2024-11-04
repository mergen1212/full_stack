package models

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Img  string `json:"img"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Img      string `json:"img"`
	Email    string `json:"email"`
	HashPass string `json:"hash_password"`
}

type UserAuth struct {
	Email    string `json:"email"`
	HashPass string `json:"hash_password"`
}

type UserReg struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Pass  string `json:"password"`
}
