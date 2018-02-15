package main

import (
	"controllers"
	"fmt"
	"log"
	"os"
	"shared"
	"time"

	"github.com/globalsign/mgo"
	"github.com/joho/godotenv"
)

func loadConfigs() {
	// local
	configFiles, err := shared.GetFileNamesFromDir("./config/", ".conf", "")
	// GCP
	// configFiles, err := shared.GetFileNamesFromDir("/home/haowei/crawler/config/", ".conf", "")
	if err != nil {
		fmt.Println("設定檔讀取失敗 = ", err)
	}

	err = godotenv.Load(configFiles...)
	if err != nil {
		log.Fatal("設定檔讀取失敗 = ", err)
	}
}

func main() {
	loadConfigs()
	crawlerCtrl := controllers.NewCrawlerController()
	items := crawlerCtrl.Craw()

	mgoCtrl := controllers.NewMongoController(getSession())
	mgoCtrl.InsertItems(items)
	// items := mgoCtrl.GetItems()

	defer mgoCtrl.Session.Close()
}

func getSession() *mgo.Session {

	// Connect to our local mongo
	info := &mgo.DialInfo{
		// Addrs:    []string{"mongodb://localhost:27017"},
		Addrs:    []string{os.Getenv("MONGO_HOST")},
		Timeout:  60 * time.Second,
		Database: os.Getenv("MONGO_DATABASE"),
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
	}

	// s, err := mgo.Dial("mongodb://localhost")
	s, err := mgo.DialWithInfo(info)

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
