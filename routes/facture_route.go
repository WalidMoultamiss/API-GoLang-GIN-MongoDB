package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.POST("/facture", controllers.CreateFacture())
	router.GET("/facturemany/:i", controllers.CreateManyFactures())
	router.GET("/facture", controllers.GetAllNotPaidFactures())
	router.GET("/facture/all", controllers.GetAllFactures())
	router.GET("/facture/:serial", controllers.GetFactureOne())
	router.GET("/facture/pay/:serial", controllers.PayFacture())

	router.POST("/car", controllers.CreateCar())
	router.GET("/carmany/:i", controllers.CreateManyCars())
	router.GET("/car", controllers.GetAllNotPaidCars())
	router.GET("/car/all", controllers.GetAllCars())
	router.GET("/car/:matriculation", controllers.GetCarOne())
	router.GET("/car/pay/:matriculation", controllers.PayCar())

	router.NoRoute(controllers.NoRoute())

	router.GET("/", controllers.HelloWorld())

}
