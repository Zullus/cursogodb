package main

import (
	r "github.com/Zullus/usuarios/routers" //r é o nome do pacote

	"github.com/labstack/echo/middleware"
)

func main() {

	//r.App.Start(":3000")

	e := r.App

	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":3000"))

}
