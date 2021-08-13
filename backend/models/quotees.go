package models

// Quotees contains all fields for a quote
type Quotees struct {
	QuoteeId			int		`db:"quotee_id" json:"quoteeId"`
	QuotebookCollection	string	`db:"quote_book_collection" json:"quotebookCollection"`
	QuoteBookTitle		string	`db:"quote_book_name" json:"quotebookName"`
	Quotee				string	`db:"quotee_name" json:"quoteeName"`
}
