package services

import (
	"github.com/techotron/online-quote-book/backend/db"
	"github.com/techotron/online-quote-book/backend/models"
)

// AddQuoteBook adds a new row in the quote_books table
func AddQuoteBook(quoteBook string) (err error) {
	_, err = db.Conn.Exec(`INSERT INTO quote_books(quote_book_name, created_on, last_updated) VALUES ($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, quoteBook)
	return err
}

// GetAllQuoteBooks returns all quote books from the quote_books table
func GetAllQuoteBooks() (quoteBooks []models.QuoteBooks, err error) {
	rows, err := db.Conn.Queryx("SELECT * FROM quote_books ORDER BY quote_book_name ASC")
	if err != nil {
		return quoteBooks, err
	}
	for rows.Next() {
		quoteBook := models.QuoteBooks{}
		err = rows.StructScan(&quoteBook)
		if err != nil {
			return quoteBooks, err
		}
		quoteBooks = append(quoteBooks, quoteBook)
	}
	return quoteBooks, err
}
