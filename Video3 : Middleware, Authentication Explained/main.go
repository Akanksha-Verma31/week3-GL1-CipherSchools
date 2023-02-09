package main

import (
	"fmt"
	"log"

	"go-application/database"
	"go-application/handler"
	"go-application/routers"

	"github.com/gin-gonic/gin"
)

func init() { //1st
	database.Setup() //establish the database connection
}
func init() { //2nd
	fmt.Print(1)
}
func init() { //3nd
	fmt.Print(2)
}

func main() {
	engine := gin.Default() //get the default engine for further customizatikon
	api := handler.Handler{
		DB: database.GetDB(), //set the handler DB
	} //get the customized engine

	routers.BookRouter(engine, api)

	err := engine.Run("127.0.0.1:8080") //start the engine
	if err != nil {
		log.Fatal(err)
	}
}