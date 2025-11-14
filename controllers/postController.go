package controllers

import (
	"go_crud/initializers"
	"go_crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Postscreate(c *gin.Context) {
	var Body struct {
		Body  *string
		Title *string
	}
	c.Bind(&Body)

	if Body.Title == nil || Body.Body == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and Body are required"})
		return
	}

	// create post
	post := models.Post{Title: *Body.Title, Body: *Body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return JSON response
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	// response
	c.JSON(http.StatusAccepted, posts)
}

func PostsShow(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	// return the post directly
	c.JSON(http.StatusOK, post)
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	// delete post
	// initializers.DB.Delete(&models.Post{}, id)
	initializers.DB.Unscoped().Delete(&models.Post{}, id)

	// return the post directly
	c.Status(http.StatusOK)
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	var Body struct {
		Title string
		Body  string
	}
	c.Bind(&Body)

	initializers.DB.Model(&post).Updates(models.Post{Title: Body.Title, Body: Body.Body})

	c.JSON(http.StatusOK, post)
}
