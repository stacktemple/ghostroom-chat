package cronjob

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func CleanAllTables(db *sqlx.DB) {
	tables := []string{"messages", "room_guests", "rooms"}
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("DELETE FROM %s", table))
		if err != nil {
			fmt.Printf("Failed to clean table %s: %v\n", table, err)
		} else {
			fmt.Printf("Table %s cleaned successfully\n", table)
		}
	}
}
