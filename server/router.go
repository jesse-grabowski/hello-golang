package server

import (
	"com.jessegrabowski/go-webapp/business/sampling"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	sampling.ConfigureRouter(router)
	return router
}
