package routers

import (
	"github.com/emails/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("html/*")
	//router.LoadHTMLFiles("templates/index.tmpl")

	router.GET("/", controllers.EmailController)
	router.POST("/", controllers.EmailSenderController)
	return router

}
