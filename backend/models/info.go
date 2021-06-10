package models

// Info contains backend information
type Info struct {
	DBSchemaVersion		int		`db:"version" json:"dbVersion"`
	DBSchemaDirty		bool	`db:"dirty" json:"dbDirty"`
}
