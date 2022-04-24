package controllers

import (
	"context"
	"fmt"
	"gin-mongo-api/configs"
	"gin-mongo-api/models"
	"gin-mongo-api/responses"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var carCollection *mongo.Collection = configs.GetCollection(configs.DB, "cars")

func GetCarOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var car models.Car

		Matriculation := c.Param("matriculation")
		fmt.Println("Matriculation", Matriculation)

		err := carCollection.FindOne(ctx, bson.M{"matriculation": Matriculation, "status": "pending"}).Decode(&car)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FactureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.FactureResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": car}})
	}
}

func PayCar() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		matriculation := c.Param("matriculation")

		update := bson.M{
			"$set": bson.M{
				"status": "paid",
			},
		}

		result, err := carCollection.UpdateOne(ctx, bson.M{"matriculation": matriculation}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FactureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.FactureResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetAllNotPaidCars() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var cars []models.Car
		cursor, err := carCollection.Find(ctx, bson.M{"status": "pending"})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FactureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		for cursor.Next(ctx) {
			var car models.Car
			err := cursor.Decode(&car)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.FactureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			cars = append(cars, car)
		}

		c.JSON(http.StatusOK, responses.FactureResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": cars}})
	}
}
func GetAllCars() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var cars []models.Car
		cursor, err := carCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FactureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		for cursor.Next(ctx) {
			var car models.Car
			err := cursor.Decode(&car)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.FactureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			cars = append(cars, car)
		}

		c.JSON(http.StatusOK, responses.FactureResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": cars}})
	}
}

func CreateCar() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var car models.Car
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&car); err != nil {
			c.JSON(http.StatusBadRequest, responses.FactureResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&car); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.FactureResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newCar := models.Car{
			Id:            primitive.NewObjectID(),
			Matriculation: car.Matriculation,
			Year:          car.Year,
			HorsePower:    car.HorsePower,
			Gas:           car.Gas,
			Status:        "pending",
		}

		result, err := carCollection.InsertOne(ctx, newCar)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FactureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.FactureResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func CreateManyCars() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		//get number from query params
		number, err := strconv.Atoi(c.Param("i"))

		if err != nil {
			c.JSON(http.StatusBadRequest, responses.FactureResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//create many cars randomly
		cars := []interface{}{}

		//array of strings
		gas := []string{"diesel", "essence"}

		//random comapany

		for i := 0; i < number; i++ {

			//random letter from A to Z
			randomLetter := string(rune(rand.Intn(26) + 65))

			//create random matriculation
			randomMatriculation := fmt.Sprintf("%d-%s-%d", rand.Intn(1000), randomLetter, rand.Intn(99))

			//random year between 2000 and 2022
			randomYear := rand.Intn(2222-2000) + 2000

			randomHorsePower := rand.Intn(11-6) + 6

			facture := models.Car{
				Id:            primitive.NewObjectID(),
				Matriculation: randomMatriculation,
				Year:          randomYear,
				HorsePower:    randomHorsePower,
				Gas:           gas[rand.Intn(len(gas))],
				Status:        "pending",
			}

			cars = append(cars, facture)
		}

		result, err := carCollection.InsertMany(ctx, cars)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FactureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.FactureResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})

	}
}
