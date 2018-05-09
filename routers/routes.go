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

	App.GET("/add", controllers.Add)

	App.GET("/atualizar/:id", controllers.Update)

	api := App.Group("/v1/")

	api.POST("insert", controllers.Inserir)

	api.GET("insert", controllers.InserGet)

	api.DELETE("delete/:id", controllers.Deletar)

	api.PUT("update/:id", controllers.Atualizar)

}
