package controllers

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/Zullus/usuarios/models"
)

//Home é a página inicial da aplicação
func Home(c echo.Context) error {

	return c.String(http.StatusOK, "Olá Mundo!")

}

//InserGet fds
func InserGet(c echo.Context) error {

	return c.String(http.StatusOK, "Chegamos no Insere!")

}

//Inserir é a função que insere no banco de dados
func Inserir(c echo.Context) error {

	nome := c.FormValue("nome")
	email := c.FormValue("email")

	var ususario models.Usuario
	ususario.Nome = nome
	ususario.Email = email

	if nome != "" && email != "" {

		if _, err := models.UsuarioModel.Insert(ususario); err != nil {

			return c.JSON(http.StatusBadRequest, map[string]string{

				"mensagem": "Não foi possível adicionar o registro no banco!",
			})

		}

		return c.JSON(http.StatusCreated, map[string]string{

			"mensagem": "Os dados foram inseridos no banco de dados com sucesso!",
		})

	}

	return c.JSON(http.StatusBadRequest, map[string]string{

		"mensagem": "Os campos não podem ser vazios!",
	})
}
