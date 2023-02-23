package app

import (
	"github.com/fajarcandraaa/go-mux-crud/entity/userentity"
)

// SetMigrationTable is used to register entity model which want to be migrate
func SetMigrationTable() []interface{} {
	var migrationData = []interface{}{
		&userentity.User{},
	}

	return migrationData
}
