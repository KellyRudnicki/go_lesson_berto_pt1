package main

import "github.com/gin-gonic/gin"

// defining a puppy here, aka the object/class go version is a struct
// structs are objects/classes 

type Puppy struct {
	Name string
	Breed string
	Age int 
	Male bool 
}

func handlePuppy(context *gin.Context) {
	puppy := getPuppy()
	context.JSON(200, puppy)
}

func getPuppy() Puppy {
	return Puppy{
		Name: "Bixby",
		Breed: "Retriever",
	}
}
