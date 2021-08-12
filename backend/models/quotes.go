package models

// Quotes contains all fields for a quote
type Quotes struct {
	QuotebookCollection	string	`db:"quote_book_collection" json:"quotebookCollection"`
	QuoteBookTitle		string	`db:"quote_book_name" json:"quotebookName"`
	QuoteText			string	`db:"quote_text" json:"quoteText"`
	Quotee				string	`db:"quotee_name" json:"quotee"`
	Witness				string	`db:"witness_name" json:"witness"`
	QuoteDate			string	`db:"quote_date" json:"quoteDate"`
}
