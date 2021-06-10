package services

import (
	"github.com/techotron/online-quote-book/backend/db"
	"github.com/techotron/online-quote-book/backend/models"
)

// GetQuotes returns quotes from the database for a specified quote book
func GetQuotes(quoteBook string) (quotes []models.Quotes, sqlError error) {
	quotes = []models.Quotes{}
	rows, sqlError := db.Conn.Queryx(`SELECT 
			quote_book_name,
			quote_text,
			quotee_name,
			witness_name,
			quote_date
		FROM quotes 
		LEFT JOIN quote_books ON quotes.quote_book_id = quote_books.quote_book_id
		LEFT JOIN quotees ON quotes.quotee_id = quotees.quotee_id
		LEFT JOIN witnesses ON quotes.witness_id = witnesses.witness_id
		WHERE quote_books.quote_book_name=$1`, quoteBook)
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
