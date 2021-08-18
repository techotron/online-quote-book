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
			quotes.quote_text,
			quotes.quotee_name,
			quotes.witness_name,
			quotes.quote_date
		FROM quotes
		LEFT JOIN quote_books ON quotes.quote_book_collection = quote_books.quote_book_collection AND quotes.quote_book_name = quote_books.quote_book_name
		LEFT JOIN quotees ON quotes.quotee_name = quotees.quotee_name
		LEFT JOIN witnesses ON quotes.witness_name = witnesses.witness_name
		WHERE quote_books.quote_book_collection=$1 AND quote_books.quote_book_name=$2
		ORDER BY quotes.quote_date ASC`, quotebookCollection, quotebook)
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

// AddQuote add a new quote to the quotes table
func AddQuote(q models.Quotes) (err error) {
	_, err = db.Conn.Exec(`INSERT INTO quotes(
		quote_book_collection,
		quote_book_name,
		quote_text,
		quotee_name,
		witness_name,
		is_deleted,
		quote_date,
		inserted_date) VALUES (
			$1, 
			$2, 
			$3, 
			$4,
			$5,
			FALSE,
			$6,
			CURRENT_TIMESTAMP
		)`, q.QuotebookCollection, q.QuoteBookTitle, q.QuoteText, q.Quotee, q.Witness, q.QuoteDate)
	return err
}
