package controllers

import "github.com/globalsign/mgo"

// MongoController mongo
type MongoController struct {
	session *mgo.Session
}

// NewMongoController 建立MongoController
func NewMongoController(s *mgo.Session) *MongoController {
	return &MongoController{s}
}
