package main

import "fmt"
import "github.com/gin-gonic/gin"

// the last word of the import file is used below in the function when exporting information, the D is capital because it's exported and it's public

func main() {
	groceryList := []string{"potato", "tomato"}
	for _, grocery := range groceryList {
		fmt.Println("Grocery Item: %s", grocery)
	}

	router := gin.Default()
	
	router.GET("/puppy", handlePuppy)
	router.Run()
}

// the := can only be done within the function, it's used for creating a new variable
// capital D means its exported (public method)
// go likes single letter variables (changed r to router)

// to clean up the code, we created another file with the puppydetails 