package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableTest, downCreateTableTest)
}

func upCreateTableTest(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS test_table (
		id varchar primary key,
		order_id varchar,
		customer_id varchar,
		created_at timestamp NOT NULL,
		updated_at timestamp NOT NULL
	);`
	if _, err := tx.Exec(query); err != nil {
		return err
	}
	return nil
}

func downCreateTableTest(tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS test_table;`
	if _, err := tx.Exec(query); err != nil {
		return err
	}
	return nil
}
