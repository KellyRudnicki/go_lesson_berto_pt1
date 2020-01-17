package main

import "github.com/gin-gonic/gin"
// the last word of the import file is used below in the function when exporting information, the D is capital because it's exported and it's public

func main() {
	router := gin.Default()
	// the := can only be done within the function, it's used for creating a new variable
	// capital D means its exported (public method)
	// go likes single letter variables (changed r to router)

	router.GET("/ping", func(context *gin.Context) {
		// above says were sending a 
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}