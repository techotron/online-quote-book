package models

// Info contains backend information
type Info struct {
	DBSchemaVersion		int		`db:"version" json:"db_version"`
	DBSchemaDirty		bool	`db:"dirty" json:"db_dirty"`
}
