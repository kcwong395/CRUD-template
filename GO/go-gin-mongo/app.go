package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kcwong395/go-gin-mongo/api"
	"github.com/kcwong395/go-gin-mongo/dbUtil"
	"log"
)

func main() {

	dbw, err := dbUtil.Init()
	if err != nil {
		panic(err)
	}
	defer dbw.Close()

	// init gin
	engine := gin.Default()

	// define the route here
	personApi := api.PersonApi{DBWrapper: dbw}
	engine.GET("/people", personApi.GetPeople)
	engine.POST("/people", personApi.AddPerson)

	// Start the go server
	if err := engine.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
