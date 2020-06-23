package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EmailController(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "GoSpike",
	})
}
