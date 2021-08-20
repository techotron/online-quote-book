package services

import (
	"github.com/techotron/online-quote-book/backend/db"
	"github.com/techotron/online-quote-book/backend/models"
)

// GetQuotees returns quotees from the database for a specified quote book
func GetQuotees(quotebookCollection, quotebook string) (quotees []models.Quotees, err error) {
	rows, err := db.Conn.Queryx(`SELECT * FROM quotees WHERE quote_book_collection=$1 AND quote_book_name=$2`, quotebookCollection, quotebook)
	if err != nil {
		return quotees, err
	}
	for rows.Next() {
		q := models.Quotees{}
		err = rows.StructScan(&q)
		if err != nil {
			return quotees, err
		}
		quotees = append(quotees, q)
	}
	return quotees, err
}

// GetQuotee returns the quotee row given the quotee name, book and collection
func GetQuotee(quoteeName, quotebookCollection, quotebook string) (quotee models.Quotees, err error) {
	err = db.Conn.Get(&quotee, `SELECT * FROM quotees WHERE quote_book_collection=$1 AND quote_book_name=$2 AND quotee_name=$3`, quotebookCollection, quotebook, quoteeName)
	return quotee, err
}

// AddQuotee adds a quotee row to the quotees table
func AddQuotee(q models.Quotees) (err error) {
	_, err = db.Conn.Exec(`INSERT INTO quotees(quote_book_collection, quote_book_name, quotee_name) 
		VALUES ($1, $2, $3)
		ON CONFLICT (quote_book_collection, quote_book_name, quotee_name) DO NOTHING`, q.QuotebookCollection, q.QuoteBookTitle, q.Quotee)
	return err
}