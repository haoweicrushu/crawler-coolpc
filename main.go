package main

import (
	"controllers"
	"time"

	"github.com/globalsign/mgo"
)

func main() {
	crawlerCtrl := controllers.NewCrawlerController()
	crawlerCtrl.Crawler()
}

func getSession() *mgo.Session {

	// Connect to our local mongo

	info := &mgo.DialInfo{
		Addrs:    []string{"mongodb://localhost"},
		Timeout:  60 * time.Second,
		Database: "",
		Username: "",
		Password: ""
	}

	// s, err := mgo.Dial("mongodb://localhost")
	s, err := mgo.DialWithInfo(info)

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
