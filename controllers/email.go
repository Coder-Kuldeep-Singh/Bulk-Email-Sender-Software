package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EmailController(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "GoSpike",
	})
}

func EmailSenderController(c *gin.Context) {
	_, err := c.MultipartForm()
	if err != nil {
		log.Println(err.Error())
		c.HTML(http.StatusOK, "index.html", gin.H{
			"formError": "Error to Parse Form",
		})
	}
}
