package controllers

import (
	"net/http"

	"github.com/techotron/online-quote-book/backend/constants"
	log "github.com/techotron/online-quote-book/backend/log"
	"github.com/techotron/online-quote-book/backend/services"

	"github.com/gin-gonic/gin"
)

// GetQuotes gets all quotes from a specified quote book
func GetQuotes(c *gin.Context) {
	quotebook := c.Param("quotebook")
	quotebookCollection := c.Param("quotebookCollection")
	q, err := services.GetQuotes(quotebookCollection, quotebook)
	if err != nil {
		if err.Error() == constants.ErrorNoRowsFound {
			log.Warnf("No quotes found for collection: %s quote book: %s", quotebookCollection, quotebook)
			c.JSON(http.StatusNotFound, q)
			return
		}
		log.Errorf("Server error: %s", err)
		c.JSON(http.StatusInternalServerError, MessageHandler(constants.MessageInternalError))
		return
	}
	c.JSON(http.StatusOK, q)
}

// AddQuote creates a new quotebook in the quote_book table
func AddQuote(c *gin.Context) {
	quotebook := c.Param("quotebook")
	quotebookCollection := c.Param("quotebookCollection")

	err := services.AddQuote(quotebookCollection, quotebook, "This is a new quote from the api", "test quotee", "test witness")
	if err != nil {
		log.Errorf("Failed to add quote: %s/%s. Server error: %s", quotebookCollection, quotebook, err)
		c.JSON(http.StatusInternalServerError, MessageHandler(constants.MessageInternalError))
		return
	}
	log.Infof("Added new quote in %s/%s", quotebookCollection, quotebook)
	c.JSON(http.StatusOK, MessageHandler("TODO: add me as constant"))
}