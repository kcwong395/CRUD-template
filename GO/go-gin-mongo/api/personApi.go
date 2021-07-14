package api

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kcwong395/go-gin-mongo/dbUtil"
	"github.com/kcwong395/go-gin-mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
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

func (p *PersonApi) UpdatePersonById(ginContext *gin.Context) {
	peopleCollection := p.DBWrapper.DB.Collection("people")

	id, _ := primitive.ObjectIDFromHex(ginContext.Param("id"))

	person := model.Person{}
	err := ginContext.ShouldBind(&person)
	if err == nil {
		result, _ := peopleCollection.ReplaceOne(context.Background(), bson.M{"_id": id}, person)
		if result != nil {
			ginContext.IndentedJSON(200, result)
		} else {
			_ = ginContext.AbortWithError(404, errors.New("person does not exist"))
		}
	} else {
		_ = ginContext.AbortWithError(404, errors.New("input parameters are problematic"))
	}
}

func (p *PersonApi) DeletePersonById(ginContext *gin.Context) {
	peopleCollection := p.DBWrapper.DB.Collection("people")

	id, _ := primitive.ObjectIDFromHex(ginContext.Param("id"))

	_, err := peopleCollection.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}})

	if err == nil {
		ginContext.IndentedJSON(200, http.StatusOK)
	} else {
		_ = ginContext.AbortWithError(404, errors.New("target person does not exist"))
	}
}
