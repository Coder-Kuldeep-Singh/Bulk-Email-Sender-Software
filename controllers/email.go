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
	// fmt.Printf("to_email :%s\nfrom_email : %s\n", to_email, from_email)
	if err != nil {
		log.Println(err.Error())
		c.HTML(http.StatusOK, "index.html", gin.H{
			"formError": "Error to Parse Form",
		})
	} else {
		to_email := c.Request.FormValue("to_email")
		from_email := c.Request.FormValue("from_email")
		password := c.Request.FormValue("password")
		message := c.Request.FormValue("message")
		if to_email == "" {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"formError": "To email is Blank",
			})
		} else if from_email == "" {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"formError": "Form email is blank",
			})
		} else if password == "" {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"formError": "Password is Blank",
			})
		} else if message == "" {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"formError": "Message is Blank",
			})
		} else if to_email != "" || from_email != "" || password != "" || message != "" {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"Success": "You're mail sent",
			})
		} else {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"fromError": "Maximum fields are blank",
			})
		}

	}
}
