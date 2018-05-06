package routers

import (
	"github.com/Zullus/usuarios/controllers"
	"github.com/labstack/echo"
)

//App a aplicação de rotas em echo
var App *echo.Echo

func init() {

	App = echo.New()

	//Rota para a pagina inical
	App.GET("/", controllers.Home)

	api := App.Group("/v1/")

	api.POST("insert", controllers.Inserir)

	api.GET("insert", controllers.InserGet)

	api.DELETE("delete/:id", controllers.Deletar)
}
