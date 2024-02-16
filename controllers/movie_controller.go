package controllers

import "github.com/gin-gonic/gin"

func ShowMovies(c *gin.Context) {
	c.JSON(200, gin.H{
		"values": "ok",
	})
}
