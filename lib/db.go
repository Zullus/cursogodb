package lib

import (
	"log"

	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var configuracao = mysql.ConnectionURL{
	Host:     "192.168.0.10",
	User:     "cursogo",
	Password: "nx74205",
	Database: "cursogo",
}

//Sess é a variavel que faz aconexão com o banco de dados
var Sess db.Database

func init() {

	var err error

	Sess, err = mysql.Open(configuracao)

	if err != nil {

		log.Fatal(err.Error())
	}
}
