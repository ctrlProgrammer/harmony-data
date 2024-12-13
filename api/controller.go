package api

import (
	"auth/api/database"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type API struct {
	Router
	state bool
}

// SECTION - Data

func (api *API) getSellers(c *gin.Context) {
	results, err := database.GetSellers(api.database)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "data": results})
}

func (api *API) getSellersByDistrict(c *gin.Context) {
	results, err := database.GetSellersByDistrict(api.database, c.Param("city"), c.Param("district"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "data": results})
}

func (api *API) getSellersByCity(c *gin.Context) {
	results, err := database.GetSellersByCity(api.database, c.Param("city"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "data": results})
}

func (api *API) getDistricts(c *gin.Context) {
	results, err := database.GetDistricts(api.database)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "data": results})
}

func (api *API) getDistrictsByCity(c *gin.Context) {
	results, err := database.GetDistrictsByCity(api.database, c.Param("city"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "data": results})
}

func (api *API) createRoutes() {
	// API - GET - /status
	// return the state of the API, if it fails internally must returns false
	api.ValidatedGet("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"state": api.state})
	})

	// API - GET - /sellers
	// This methods will use front and back cache to reduce the database load based on request
	api.ValidatedGet("/sellers", api.getSellers)

	// API - GET - /sellers-by-district/:district
	// This methods will use front and back cache to reduce the database load based on request
	api.ValidatedGet("/sellers-by-district/:city/:district", api.getSellersByDistrict)

	// API - GET - /sellers-by-city/:city
	// This methods will use front and back cache to reduce the database load based on request
	api.ValidatedGet("/sellers-by-city/:city", api.getSellersByCity)

	// API - GET - /districts
	// The information will be considered static so it will use a simple cache method
	api.ValidatedGet("/districts", api.getDistricts)

	// API - GET - /districts/:city
	// The information will be considered static so it will use a simple cache method
	api.ValidatedGet("/districts/:city", api.getDistrictsByCity)

	api.logger.Info("Initilizing routes")
}

func (api *API) Initialize(logger *zap.SugaredLogger, database *mongo.Database) {
	api.logger = logger
	api.database = database

	api.state = true
	api.logs = true

	api.logger.Info("Created sessions map")

	api.gin = gin.Default()
	api.gin.Use(CorsMiddleware())

	api.createRoutes()

	api.gin.Run(":" + os.Getenv("PORT"))

}
