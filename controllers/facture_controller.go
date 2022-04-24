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
	"github.com/go-playground/validator/v10"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var factureCollection *mongo.Collection = configs.GetCollection(configs.DB, "factures")
var validate = validator.New()

func GetFactureOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var factures models.Facture

		serial := c.Param("serial")
		fmt.Println("serial", serial)

		err := factureCollection.FindOne(ctx, bson.M{"serial": serial, "status": "pending"}).Decode(&factures)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FactureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.FactureResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": factures}})
	}
}

func PayFacture() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		serial := c.Param("serial")

		update := bson.M{
			"$set": bson.M{
				"status": "paid",
			},
		}

		result, err := factureCollection.UpdateOne(ctx, bson.M{"serial": serial}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FactureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.FactureResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetAllNotPaidFactures() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var factures []models.Facture
		cursor, err := factureCollection.Find(ctx, bson.M{"status": "pending"})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FactureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		for cursor.Next(ctx) {
			var facture models.Facture
			err := cursor.Decode(&facture)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.FactureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			factures = append(factures, facture)
		}

		c.JSON(http.StatusOK, responses.FactureResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": factures}})
	}
}
func GetAllFactures() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var factures []models.Facture
		cursor, err := factureCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FactureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		for cursor.Next(ctx) {
			var facture models.Facture
			err := cursor.Decode(&facture)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.FactureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			factures = append(factures, facture)
		}

		c.JSON(http.StatusOK, responses.FactureResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": factures}})
	}
}

func CreateFacture() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var facture models.Facture
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&facture); err != nil {
			c.JSON(http.StatusBadRequest, responses.FactureResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&facture); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.FactureResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newFacture := models.Facture{
			Id:      primitive.NewObjectID(),
			Serial:  facture.Serial,
			Price:   facture.Price,
			Company: facture.Company,
			Status:  facture.Status,
		}

		result, err := factureCollection.InsertOne(ctx, newFacture)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FactureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.FactureResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func CreateManyFactures() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		//get number from query params
		number, err := strconv.Atoi(c.Param("i"))

		if err != nil {
			c.JSON(http.StatusBadRequest, responses.FactureResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//create many factures randomly
		factures := []interface{}{}

		//array of strings
		companies := []string{"ONE", "LYDEC", "RADES"}

		//random comapany

		for i := 0; i < number; i++ {
			randomCompany := companies[rand.Intn(len(companies))]
			facture := models.Facture{
				Id: primitive.NewObjectID(),
				//random serial number 10 digits as string
				Serial:  strconv.Itoa(rand.Intn(9999999999)),
				Price:   rand.Intn(100),
				Company: randomCompany,
				Status:  "pending",
			}

			factures = append(factures, facture)
		}

		result, err := factureCollection.InsertMany(ctx, factures)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FactureResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.FactureResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})

	}
}
