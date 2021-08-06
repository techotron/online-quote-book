package controllers

import (
	"net/http"

	"github.com/techotron/online-quote-book/backend/constants"
	"github.com/techotron/online-quote-book/backend/log"
	"github.com/techotron/online-quote-book/backend/services"

	"github.com/gin-gonic/gin"
)

// AddQuoteBook creates a new quotebook in the quote_book table
func AddQuoteBook(c *gin.Context) {
	quoteBook := c.Param("quoteBook")
	err := services.AddQuoteBook(quoteBook)
	if err != nil {
		log.Errorf("Failed to add quotebook: %s. Server error: %s", quoteBook, err)
		c.JSON(http.StatusInternalServerError, MessageHandler(constants.MessageInternalError))
		return
	}
	log.Infof("Added new quote book: %s", quoteBook)
	c.JSON(http.StatusOK, MessageHandler(constants.NewQuoteBookAddedSuccess))
}

// GetAllQuoteBooks returns all quote books from the database
func GetAllQuoteBooks(c *gin.Context) {
	qbs, err := services.GetAllQuoteBooks()
	if err != nil {
		log.Errorf("Failed to return all quote books. Error: %s", err)
		c.JSON(http.StatusInternalServerError, MessageHandler(constants.MessageInternalError))
		return
	}
	log.Infof("Returned all quotebooks")
	c.JSON(http.StatusOK, qbs)
}
