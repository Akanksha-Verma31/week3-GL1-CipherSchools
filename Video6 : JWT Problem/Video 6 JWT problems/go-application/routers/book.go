package routers

import (
	"go-application/handler"

	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.Engine, api handler.Handler) {
	//get the default engine for further customizatikon
	// api := handler.Handler{
	// 	DB: database.GetDB(), //set the handler DB
	// }

	router.GET("/books", api.GetBooks) //set the function for http://localhost:8080/books : Get request
	//while calling handler function, gin will pass *gin.Context in the handler function
	router.POST("/books", api.SaveBook)
	router.GET("/books/:id", api.GetBookByID)
	router.DELETE("/books/:id", api.DeleteBookByID)
	router.PUT("/books", api.UpdateBookByID)
	//	router.DELETE("/books/:id", api.DeleteBook)
	// return router
}

/*
http://localhost:8080/book/11         - params
http://localhost:8080/book?id=11&name=abc	  - query


*/
