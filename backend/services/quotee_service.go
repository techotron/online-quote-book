package services

import (
	"github.com/techotron/online-quote-book/backend/db"
	"github.com/techotron/online-quote-book/backend/models"
)

// GetQuotees returns quotees from the database for a specified quote book
func GetQuotees(quotebookCollection, quotebook string) (quotees []models.Quotees, sqlError error) {
	quotees = []models.Quotees{}
	rows, sqlError := db.Conn.Queryx(`SELECT * FROM quotees WHERE quote_book_collection=$1 AND quote_book_name=$2`, quotebookCollection, quotebook)
	if sqlError != nil {
		return quotees, sqlError
	}
	for rows.Next() {
		q := models.Quotees{}
		sqlError = rows.StructScan(&q)
		if sqlError != nil {
			return quotees, sqlError
		}
		quotees = append(quotees, q)
	}
	return quotees, sqlError
}

// GetQuotee returns the quotee row given the name, book and collection