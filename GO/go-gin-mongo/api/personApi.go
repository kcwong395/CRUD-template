package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/kcwong395/go-gin-mongo/dbUtil"
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
