package models

import "time"

// QuoteBooks contains all fields for a quote book
type QuoteBooks struct {
	QuoteBookId		int			`db:"quote_book_id" json:"quoteBookId"`
	QuoteBookTitle	string		`db:"quote_book_name" json:"quoteBookName"`
	CreatedOn		time.Time	`db:"created_on" json:"createdOn"`
	LastUpdated		time.Time	`db:"last_updated" json:"lastUpdated"`
}
