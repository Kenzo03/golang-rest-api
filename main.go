package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type resp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Kale", Price: 59.25},
	{ID: "2", Title: "Green Train", Artist: "Mayer Kale", Price: 25.25},
	{ID: "3", Title: "Red Train", Artist: "Tuly San", Price: 30},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if newAlbum.ID == "" {
		// var newResp = resp
		c.IndentedJSON(http.StatusBadRequest, resp{Status: "error", Message: "ID can't be empty!"})
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
