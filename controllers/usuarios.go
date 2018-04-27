package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

//Home é a página inicial da aplicação
func Home(c echo.Context) error {

	return c.String(http.StatusOK, "Olá Mundo!")

}
