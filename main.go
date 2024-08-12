package main

import (
	"crudMongo/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/", controllers.GreetUser)
	r.POST("/create", controllers.CreateWatchList)
	r.POST("/multiCreate", controllers.CreateMultipleWatchList)
	r.GET("/watchlist", controllers.Getwatchlist)
	r.GET("/deletewatchlist", controllers.Deletewatchlist)
	r.Run()
}
