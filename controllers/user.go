package controllers

import (
	"crudMongo/database"
	"crudMongo/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GreetUser(c *gin.Context) {

	name := c.Request.FormValue("name")

	responseString := fmt.Sprintf("Hello va da venaa mavane  %v", name)

	c.JSON(http.StatusOK, gin.H{"welcome": responseString})
}

func CreateWatchList(c *gin.Context) {
	var newMovie models.Netflix

	if c.Request.FormValue("movName") != "" {
		newMovie.Movie = c.Request.FormValue("movName")
	}

	if c.Request.FormValue("watched") != "" {
		watched := c.Request.FormValue("watched")

		if watched == "0" {
			newMovie.Watched = false
		} else if watched == "1" {
			newMovie.Watched = true
		}
	}

	created, err := database.CreateMovieWatchList(newMovie)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"Successfully created the watchlist": created})

}

func CreateMultipleWatchList(c *gin.Context) {

	file, fileHeader, err := c.Request.FormFile("log")
	if err != nil {
		panic(err)
	}

	fileData, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var documents []interface{}

	err = json.Unmarshal(fileData, &documents)
	if err != nil {
		panic(err)
	}

	result, err := database.CreateMultipleWatchList(documents)
	if err != nil {
		panic(err)
	} else if result {
		c.JSON(http.StatusOK, gin.H{"Successfully read the file Data": fileHeader.Filename})
	}
}

func Getwatchlist(c *gin.Context) {

	data, err := database.GetWatchList()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"Data": data})
}

func Deletewatchlist(c *gin.Context) {

	id := c.Request.URL.Query().Get("id")
	if id != "" {
		err := database.DeleteWatchList(id)
		if err != nil {
			log.Fatal(err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"Success": "Deleted the watchlist successfully"})
}
