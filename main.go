package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/routes" //add this

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	//routes
	routes.Routes(router) //add this

	router.Run("localhost:6000")
}
