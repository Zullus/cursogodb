package models

//Usuario representa o usuário na base de dados
type Usuario struct {
	ID    int    `json:"id" db:"id"`
	Nome  string `json:"nome" db:"nome"`
	Email string `json:"email" db:"email"`
}
