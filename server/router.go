package server

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.New()

	return router
}
