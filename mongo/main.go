package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"github.com/gin-gonic/gin"
)
// the structure of Daily News
type News struct {
	Title   string   `bson:"title"`
	Content string   `bson:"content"`
	By 		string	 `bson:"by"`
	Url 	string   `bson:"url"`
	CreatedAt string `bson:"created_at"`
}
// setting up router groups
func setupRouter(r *gin.Engine)  {
	r.GET("/news", getNews)
}
// initialize mongodb
func mgoInit() *mgo.Session {
	mgo, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	return mgo
}
// get json data from daily news
var news []News
var dbName string
var collectionNews string
func getNews(c *gin.Context) {
	mgo := mgoInit()
	collection := mgo.DB(dbName).C(collectionNews)
	iter := collection.Find(nil).Iter()
	err    := iter.All(&news)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	defer mgo.Close()
	defer iter.Close()
	c.JSON(200, gin.H{
		"news": news,
	})
}
// start the program
func main() {
	r := gin.Default()
	setupRouter(r)
	r.Run(":9090")
}