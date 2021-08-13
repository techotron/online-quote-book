package services

import (
	"github.com/techotron/online-quote-book/backend/db"
	"github.com/techotron/online-quote-book/backend/models"
)

// GetWitnesses returns witnesses from the database for a specified quote book
func GetWitnesses(quotebookCollection, quotebook string) (witnesses []models.Witnesses, sqlError error) {
	witnesses = []models.Witnesses{}
	rows, sqlError := db.Conn.Queryx(`SELECT * FROM witnesses WHERE quote_book_collection=$1 AND quote_book_name=$2`, quotebookCollection, quotebook)
	if sqlError != nil {
		return witnesses, sqlError
	}
	for rows.Next() {
		w := models.Witnesses{}
		sqlError = rows.StructScan(&w)
		if sqlError != nil {
			return witnesses, sqlError
		}
		witnesses = append(witnesses, w)
	}
	return witnesses, sqlError
}
