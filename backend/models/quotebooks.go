package models

import "time"

// QuoteBooks contains all fields for a quote book
type QuoteBooks struct {
	QuoteBookId    		int       `db:"quote_book_id" json:"quotebookId"`
	QuotebookCollection	string	`db:"quote_book_collection" json:"quotebookCollection"`
	QuoteBookTitle 		string    `db:"quote_book_name" json:"quotebookName"`
	CreatedOn      		time.Time `db:"created_on" json:"createdOn"`
	LastUpdated    		time.Time `db:"last_updated" json:"lastUpdated"`
}
