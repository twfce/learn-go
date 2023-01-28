package endpoints

import (
	"net/http"
	"simple-api/structs"
    "github.com/gin-gonic/gin"
)

// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, structs.Albums)
}

// PostAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
    var newAlbum structs.Album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    structs.Albums = append(structs.Albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GetAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range structs.Albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}