package controllers

import (
	"net/http"

	"github.com/techotron/online-quote-book/backend/constants"
	log "github.com/techotron/online-quote-book/backend/log"
	"github.com/techotron/online-quote-book/backend/services"

	"github.com/gin-gonic/gin"
)



// GetWitnesses gets all witnesses from a specified quote book
func GetWitnesses(c *gin.Context) {
	quotebook := c.Param("quotebook")
	quotebookCollection := c.Param("quotebookCollection")
	w, err := services.GetWitnesses(quotebookCollection, quotebook)
	if err != nil {
		if err.Error() == constants.ErrorNoRowsFound {
			log.Warnf("No witnesses found for collection: %s quote book: %s", quotebookCollection, quotebook)
			c.JSON(http.StatusNotFound, w)
			return
		}
		log.Errorf("Server error: %s", err)
		c.JSON(http.StatusInternalServerError, MessageHandler(constants.MessageInternalError))
		return
	}
	c.JSON(http.StatusOK, w)
}
