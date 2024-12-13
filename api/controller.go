package api

import (
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

}

func (api *API) createRoutes() {
	// API - GET - /status
	// return the state of the API, if it fails internally must returns false
	api.ValidatedGet("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"state": api.state})
	})

	// API - GET - /sellers
	// get all sellers
	api.ValidatedGet("/sellers", api.getSellers)

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
