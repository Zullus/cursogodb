package models

import (
	"github.com/Zullus/usuarios/lib"
)

//Usuario representa o usuário na base de dados
type Usuarios struct {
	ID    int    `json:"id" db:"id"`
	Nome  string `json:"nome" db:"nome"`
	Email string `json:"email" db:"email"`
}

//UsuarioModel recebe a tabela do banco de dados
var UsuarioModel = lib.Sess.Collection("usuarios") //temos collection e collections!
