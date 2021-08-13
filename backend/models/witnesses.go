package models

// Witnesses contains all fields for a quote
type Witnesses struct {
	WitnessId			int		`db:"witness_id" json:"witnessId"`
	QuotebookCollection	string	`db:"quote_book_collection" json:"quotebookCollection"`
	QuoteBookTitle		string	`db:"quote_book_name" json:"quotebookName"`
	Witness				string	`db:"witness_name" json:"witnessName"`
}
