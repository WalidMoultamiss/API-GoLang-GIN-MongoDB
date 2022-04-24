package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloWorld() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"menu": map[string]interface{}{
				"get all factures": map[string]interface{}{
					"url":    "/facture/all",
					"method": "GET",
				},
				"get all factures not paid": map[string]interface{}{
					"url":    "/facture",
					"method": "GET",
				},
				"create new facture": map[string]interface{}{
					"url":    "/facture",
					"method": "POST",
				},
				"Create many random factures": map[string]interface{}{
					"url":    "/facturemany/:numberOfFactures",
					"method": "GET",
				},
				"get a facture by serial and not paid": map[string]interface{}{
					"url":    "/facture/:serial",
					"method": "GET",
				},
				"get all cars": map[string]interface{}{
					"url":    "/car/all",
					"method": "GET",
				},
				"get all cars not paid": map[string]interface{}{
					"url":    "/car",
					"method": "GET",
				},
				"create new car": map[string]interface{}{
					"url":    "/car",
					"method": "POST",
				},
				"Create many random cars": map[string]interface{}{
					"url":    "/carmany/:numberOfcars",
					"method": "GET",
				},
				"get a car by matriculation and not paid": map[string]interface{}{
					"url":    "/car/:matriculation",
					"method": "GET",
				},
			},
			"schema": map[string]interface{}{
				"facture": map[string]interface{}{
					"serial":  "00000001",
					"price":   99,
					"status":  "pending",
					"company": "ONE",
				},
				"car": map[string]interface{}{
					"matriculation": "1234-A-54",
					"year":          2010,
					"horsepower":    7,
					"gas":           "diesel",
					"status":        "pending",
				},
			},
		})
	}
}

func NoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No route found",
		})
	}
}
