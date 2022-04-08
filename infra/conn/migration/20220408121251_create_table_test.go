package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableTest, downCreateTableTest)
}

func upCreateTableTest(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downCreateTableTest(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
