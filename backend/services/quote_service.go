package services

import (
	"github.com/techotron/online-quote-book/backend/db"
	"github.com/techotron/online-quote-book/backend/models"
)

// GetQuotes returns quotes from the database for a specified quote book
func GetQuotes(quotebookCollection, quotebook string) (quotes []models.Quotes, sqlError error) {
	quotes = []models.Quotes{}
	rows, sqlError := db.Conn.Queryx(`SELECT 
			quotes.quote_book_name,
			quote_text,
			quotee_name,
			witness_name,
			quote_date
		FROM quotes 
		LEFT JOIN quote_books ON quotes.quote_book_collection = quote_books.quote_book_collection AND quotes.quote_book_name = quote_books.quote_book_name
		LEFT JOIN quotees ON quotes.quotee_id = quotees.quotee_id
		LEFT JOIN witnesses ON quotes.witness_id = witnesses.witness_id
		WHERE quote_books.quote_book_collection=$1 AND quote_books.quote_book_name=$2`, quotebookCollection, quotebook)
	if sqlError != nil {
		return quotes, sqlError
	}
	for rows.Next() {
		q := models.Quotes{}
		sqlError = rows.StructScan(&q)
		if sqlError != nil {
			return quotes, sqlError
		}
		quotes = append(quotes, q)
	}
	return quotes, sqlError
}
