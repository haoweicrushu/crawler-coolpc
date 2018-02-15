package controllers

import (
	"fmt"
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

// GetItems 取得 items
func (m *MongoController) GetItems() (items []models.Item) {
	err := m.Session.DB("crawler").C("item").Find(nil).All(&items)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
