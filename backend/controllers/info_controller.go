package controllers

import (
	"net/http"

	log "github.com/techotron/online-quote-book/backend/log"
	"github.com/techotron/online-quote-book/backend/services"

	"github.com/gin-gonic/gin"
)


// GetInfo returns a basic info response to proof the server is up
func GetInfo(c *gin.Context) {
	log.Debug("GetInfo: ", c)
	i, err := services.GetSchemaInfo()
	if err != nil {
		log.Errorf("Server error: %S", err)
		c.JSON(http.StatusInternalServerError, "Failed to fetch DB schema version")
		return
	}
	c.JSON(http.StatusOK, i)
}
