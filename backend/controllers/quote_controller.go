package controllers

import (
	"net/http"

	log "github.com/techotron/online-quote-book/backend/log"
	"github.com/techotron/online-quote-book/backend/services"
	"github.com/techotron/online-quote-book/backend/constants"

	"github.com/gin-gonic/gin"
)

// GetQuotes gets all quotes from a specified quote book
func GetQuotes(c *gin.Context) {
	quoteBook := c.Param("quoteBook")
	q, err := services.GetQuotes(quoteBook)
	if err != nil {
		if err.Error() == constants.ErrorNoRowsFound {
			log.Warnf("No quotes found for quote book: %s", quoteBook)
			c.JSON(http.StatusNotFound, q)
			return
		}
		log.Errorf("Server error: %s", err)
		c.JSON(http.StatusInternalServerError, MessageHandler(constants.MessageInternalError))
		return
	}
	c.JSON(http.StatusOK, q)
}

