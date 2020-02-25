package controller

import (
	"fmt"
	"net/http"

	"github.com/tweetGo/db"

	"github.com/gin-gonic/gin"
	"github.com/tweetGo/model"
)

// TweetsIndex はインデックス画面を表示する。
func TweetsIndex(c *gin.Context) {
	var tweets []model.Tweet

	db := db.Init()
	err := db.Find(&tweets).Error
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"tweets": tweets,
	})
}

// TweetsCreate はインデックス画面を表示する。
func TweetsCreate(c *gin.Context) {

	var tweet model.Tweet
	c.ShouldBind(&tweet)
	db := db.Init()
	err := db.Create(&tweet).Error
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	}

	TweetsIndex(c)
}

// TweetsEdit はTweetの編集画面を表示する。
func TweetsEdit(c *gin.Context) {
	tweetID := c.Param("tweet_id")
	var tweet model.Tweet

	db := db.Init()
	err := db.First(&tweet, tweetID)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	}

	c.HTML(http.StatusOK, "edit.tmpl", gin.H{
		"tweet": tweet,
	})
}

// TweetsDelete はTweetを削除します。
func TweetsDelete(c *gin.Context) {
	tweetID := c.Param("tweet_id")
	var tweet model.Tweet

	db := db.Init()
	err := db.First(&tweet, tweetID)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	}

	db.Delete(&tweet)

	TweetsIndex(c)
}

// TweetsUpdate はTweetを更新します。
func TweetsUpdate(c *gin.Context) {
	tweetID := c.Param("tweet_id")
	var tweet model.Tweet
	text := c.PostForm("text")

	db := db.Init()
	err := db.First(&tweet, tweetID)
	if err != nil {
		c.AbortWithStatus(400)
		fmt.Println(err)
	}

	tweet.Text = text
	db.Save(&tweet)

	TweetsIndex(c)
}
