package services

import (
	"github.com/techotron/online-quote-book/backend/db"
	"github.com/techotron/online-quote-book/backend/models"
)

// AddQuoteBook adds a new row in the quote_books table
func AddQuoteBook(quotebookCollection, quotebook string) (err error) {
	_, err = db.Conn.Exec(`INSERT INTO quote_books(quote_book_collection, quote_book_name, created_on, last_updated) 
			VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
			ON CONFLICT (quote_book_collection, quote_book_name) DO NOTHING`, quotebookCollection, quotebook)
	return err
}

// GetAllQuoteBooks returns all quote books from the quote_books table
func GetAllQuoteBooks() (quotebooks []models.QuoteBooks, err error) {
	rows, err := db.Conn.Queryx("SELECT * FROM quote_books ORDER BY quote_book_collection ASC, quote_book_name ASC")
	if err != nil {
		return quotebooks, err
	}
	for rows.Next() {
		quotebook := models.QuoteBooks{}
		err = rows.StructScan(&quotebook)
		if err != nil {
			return quotebooks, err
		}
		quotebooks = append(quotebooks, quotebook)
	}
	return quotebooks, err
}
