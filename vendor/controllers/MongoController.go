package controllers

import (
	"models"
	"time"

	"github.com/globalsign/mgo"
)

// MongoController mongo
type MongoController struct {
	Session *mgo.Session
}

// NewMongoController 建立MongoController
func NewMongoController(s *mgo.Session) *MongoController {
	return &MongoController{s}
}

// InsertItems 把Items 寫進db
func (m *MongoController) InsertItems(items []models.Item) {
	for _, item := range items {
		item.Date = time.Now()
		m.Session.DB("crawler").C("item").Insert(item)
	}
}
