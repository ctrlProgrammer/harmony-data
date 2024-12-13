package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// Low level API functionalities
type Router struct {
	logger   *zap.SugaredLogger
	gin      *gin.Engine
	database *mongo.Database
	logs     bool
}

func (router *Router) Log(message string) {
	if router.logs {
		router.logger.Info(message)
	}
}

func (router *Router) ErrorLog(message string) {
	if router.logs {
		router.logger.Error(message)
	}
}

func (router *Router) valudateAuth(c *gin.Context) {
	token := c.GetHeader("HARMONY_MICRO_SERVICES")

	if token != os.Getenv("HARMONY_MICRO_SERVICES_KEY") {
		router.ErrorLog("Invalid micro service token key")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.Next()
}

// SECTION - Sessions
// Validate session on the auth microservice

func (api *Router) middleWareValidateSession(c *gin.Context) {
	// var request types.Logged

	// TODO Validate session on the auth microservice

	// if !success {
	// 	c.AbortWithStatus(http.StatusForbidden)
	// 	return
	// }

	c.Next()
}

// Default methods
// We can use more methods but I will reduce it to GET and POST, I know that I can use PUT, UPDATE...

func (router *Router) ValidatedGet(url string, callback func(c *gin.Context)) {
	router.gin.GET(url, router.valudateAuth, callback)
}

func (router *Router) ValidatedPost(url string, callback func(c *gin.Context)) {
	router.gin.POST(url, router.valudateAuth, callback)
}

func (router *Router) ValidatedWithSessionPost(url string, callback func(c *gin.Context)) {
	router.gin.POST(url, router.valudateAuth, router.middleWareValidateSession, callback)
}
