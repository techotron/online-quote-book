package services

import (
	"github.com/techotron/online-quote-book/backend/db"
	"github.com/techotron/online-quote-book/backend/models"
)

// GetSchemaInfo returns DB schema information from the database
func GetSchemaInfo() (models.Info, error) {
	i := models.Info{}
	err := db.Conn.Get(&i, "SELECT * FROM schema_migrations")
	return i, err
}