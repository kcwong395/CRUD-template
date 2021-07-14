package api

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kcwong395/go-gin-mongo/dbUtil"
	"github.com/kcwong395/go-gin-mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type PersonApi struct {
	DBWrapper *dbUtil.DBWrapper
}

func (p *PersonApi) GetPeople(ginContext *gin.Context) {
	peopleCollection := p.DBWrapper.DB.Collection("people")

	cursor, err := peopleCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Panic(err)
	}

	var people []bson.M
	if err := cursor.All(context.Background(), &people); err != nil {
		log.Panic(err)
	}

	ginContext.IndentedJSON(200, people)
}

func (p *PersonApi) AddPerson(ginContext *gin.Context) {
	peopleCollection := p.DBWrapper.DB.Collection("people")

	person := model.Person{}
	err := ginContext.ShouldBind(&person)
	if err != nil {
		_ = ginContext.AbortWithError(500, errors.New("input parameters are problematic"))
		return
	}

	result, err := peopleCollection.InsertOne(context.Background(), person)
	if err != nil {
		_ = ginContext.AbortWithError(500, errors.New("fail to add person"))
	} else {
		ginContext.IndentedJSON(201, result)
	}
}
