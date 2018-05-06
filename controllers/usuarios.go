package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/Zullus/usuarios/models"
)

//Home é a página inicial da aplicação
func Home(c echo.Context) error {

	var usuarios []models.Usuarios

	if err := models.UsuarioModel.Find().All(&usuarios); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{

			"mensagem": "Erro ao tentar recuperar os dados",
		})
	}

	data := map[string]interface{}{
		"titulo":   "Listagem de Usuários",
		"usuarios": usuarios,
	}

	//return c.String(http.StatusOK, "Olá Mundo!")
	return c.Render(http.StatusOK, "index.html", data)

}

//InserGet fds
func InserGet(c echo.Context) error {

	return c.String(http.StatusOK, "Chegamos no Insere!")

}

//Inserir é a função que insere no banco de dados
func Inserir(c echo.Context) error {

	nome := c.FormValue("nome")
	email := c.FormValue("email")

	var ususario models.Usuarios
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

//Deletar Usuário do banco
func Deletar(c echo.Context) error {

	usuarioID, _ := strconv.Atoi(c.Param("id"))

	resultado := models.UsuarioModel.Find("id=?", usuarioID)

	if count, _ := resultado.Count(); count < 1 {

		return c.JSON(http.StatusNotFound, map[string]string{
			"mensagem": "Não foi possível encontrar o usuário",
		})

	}

	if err := resultado.Delete(); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Não foi possíel deletar o usuário",
		})

	}

	return c.JSON(http.StatusAccepted, map[string]string{
		"mensagem": "Usuário deletado com sucesso",
	})

}
