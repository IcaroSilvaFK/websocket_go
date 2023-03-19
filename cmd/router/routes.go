package routes

import (
	"net/http"

	"github.com/IcaroSilvaFK/newsletter_go/cmd/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/ws", controllers.InitializeWS)

}
