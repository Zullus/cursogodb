package routers

import (
	"github.com/Zullus/usuarios/controllers"
	"github.com/labstack/echo"
)

//App a aplicação de rotas em echo
var App *echo.Echo

func init() {

	App = echo.New()

	App.GET("/", controllers.Home)
}
