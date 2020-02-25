package main

import (
	"github.com/gin-gonic/gin"

	"github.com/tweetGo/controller"
)

func main() {

	router := gin.Default()
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("view/*/*")
	router.GET("/tweets", controller.TweetsIndex)
	router.GET("/tweets/:tweet_id", controller.TweetsEdit)
	router.GET("/tweets/:tweet_id/delete", controller.TweetsDelete)
	router.POST("/tweets", controller.TweetsCreate)
	router.POST("/tweets/:tweet_id", controller.TweetsUpdate)
	router.Run(":8080")
}
