package main

import (
	"github.com/aabdullahgungor/golang-jwt-restapi/controller"
	"github.com/aabdullahgungor/golang-jwt-restapi/database"
	"github.com/aabdullahgungor/golang-jwt-restapi/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize Database
	database.Connect()
	database.Migrate()
	// Initialize Router
	router := initRouter()
	router.Run(":8080")

}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controller.GenerateToken)
		api.POST("/user/register", controller.RegisterUser)
		secured := api.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/ping", controller.Ping)
		}
	}
	return router
}