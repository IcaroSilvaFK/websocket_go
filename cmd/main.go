package main

import (
	"log"

	routes "github.com/IcaroSilvaFK/newsletter_go/cmd/router"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	routes.InitRoutes(app)

	log.Println("🚀Server running at port 8080")
	app.Run()
}
