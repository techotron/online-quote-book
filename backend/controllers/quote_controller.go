package controllers

import (
	"net/http"

	"github.com/techotron/online-quote-book/backend/constants"
	log "github.com/techotron/online-quote-book/backend/log"
	"github.com/techotron/online-quote-book/backend/services"
	"github.com/techotron/online-quote-book/backend/models"

	"github.com/gin-gonic/gin"
)

type postQuoteRequestBody struct {
	QuoteText	string	`json:"quoteText" binding:"required"`
	Quotee		string	`json:"quotee" binding:"required"`
	Witness		string	`json:"witness" binding:"required"`
	QuoteDate	string	`json:"quoteDate" binding:"required"`
}

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
	requestBody := postQuoteRequestBody{}
	q := models.Quotes{}

	q.QuotebookCollection = c.Param("quotebookCollection")
	q.QuoteBookTitle = c.Param("quotebook")

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Warn(err)
		c.JSON(http.StatusBadRequest, MessageHandler(constants.UnableToMarshallPayload))
		return
	}
	
	q.QuoteText = requestBody.QuoteText
	q.QuoteDate = requestBody.QuoteDate
	q.Quotee = requestBody.Quotee
	q.Witness = requestBody.Witness

	// Check if given quotee already exists in table
	quotee, err := services.GetQuotee(q.Quotee, q.QuotebookCollection, q.QuoteBookTitle)
	if err != nil {
		if err.Error() == constants.NoRowsErrorFromDB {
			quotee.QuoteBookTitle = q.QuoteBookTitle
			quotee.QuotebookCollection = q.QuotebookCollection
			quotee.Quotee = q.Quotee
			err = services.AddQuotee(quotee)
			if err != nil {
				log.Warn(err)
				c.JSON(http.StatusInternalServerError, "Failed to add new quotee")
			}
		}
		log.Warnf("Problem in querying for quotee: %s in %s/%s. Error: %s", q.Quotee, q.QuotebookCollection, q.QuoteBookTitle, err)
	}

	// Check if given witness already exists in table
	witness, err := services.GetWitness(q.Witness, q.QuotebookCollection, q.QuoteBookTitle)
	if err != nil {
		if err.Error() == constants.NoRowsErrorFromDB {
			witness.QuoteBookTitle = q.QuoteBookTitle
			witness.QuotebookCollection = q.QuotebookCollection
			witness.Witness = q.Witness
			err = services.AddWitness(witness)
			if err != nil {
				log.Warn(err)
				c.JSON(http.StatusInternalServerError, "Failed to add new witness")
			}
		}
		log.Warnf("Problem in querying for witness: %s in %s/%s. Error: %s", q.Witness, q.QuotebookCollection, q.QuoteBookTitle, err)
	}

	err = services.AddQuote(q)
	if err != nil {
		log.Errorf("Failed to add quote: %s/%s. Server error: %s", q.QuotebookCollection, q.QuoteBookTitle, err)
		c.JSON(http.StatusInternalServerError, MessageHandler(constants.MessageInternalError))
		return
	}
	log.Infof("Added new quote in %s/%s", q.QuotebookCollection, q.QuoteBookTitle)
	c.JSON(http.StatusOK, MessageHandler(constants.NewQuoteAddedSuccess))
}