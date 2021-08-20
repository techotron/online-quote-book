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

// GetWitness returns the witness row given the witness name, book and collection
func GetWitness(witnessName, quotebookCollection, quotebook string) (witness models.Witnesses, err error) {
	err = db.Conn.Get(&witness, `SELECT * FROM witnesses WHERE quote_book_collection=$1 AND quote_book_name=$2 AND witness_name=$3`, quotebookCollection, quotebook, witnessName)
	return witness, err
}

// AddWitness adds a witness row to the witnesses table
func AddWitness(w models.Witnesses) (err error) {
	_, err = db.Conn.Exec(`INSERT INTO witnesses(quote_book_collection, quote_book_name, witness_name) 
		VALUES ($1, $2, $3)
		ON CONFLICT (quote_book_collection, quote_book_name, witness_name) DO NOTHING`, w.QuotebookCollection, w.QuoteBookTitle, w.Witness)
	return err
}
