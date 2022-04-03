package sampling

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ConfigureRouter(router *gin.Engine) {
	router.POST("/entities", handleCreateEntity)
	router.GET("/entities/:id", handleReadEntity)
}

func handleCreateEntity(c *gin.Context) {
	var request Entity
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request, err = createEntity(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, request)
}

func handleReadEntity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse id"})
		return
	}

	response, err := readEntity(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
