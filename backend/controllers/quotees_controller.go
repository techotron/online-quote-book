package controllers

import (
	"net/http"

	"github.com/techotron/online-quote-book/backend/constants"
	log "github.com/techotron/online-quote-book/backend/log"
	"github.com/techotron/online-quote-book/backend/services"

	"github.com/gin-gonic/gin"
)



// GetQuotees gets all quotees from a specified quote book
func GetQuotees(c *gin.Context) {
	quotebook := c.Param("quotebook")
	quotebookCollection := c.Param("quotebookCollection")
	q, err := services.GetQuotees(quotebookCollection, quotebook)
	if err != nil {
		if err.Error() == constants.ErrorNoRowsFound {
			log.Warnf("No quotees found for collection: %s quote book: %s", quotebookCollection, quotebook)
			c.JSON(http.StatusNotFound, q)
			return
		}
		log.Errorf("Server error: %s", err)
		c.JSON(http.StatusInternalServerError, MessageHandler(constants.MessageInternalError))
		return
	}
	c.JSON(http.StatusOK, q)
}
