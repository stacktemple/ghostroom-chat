package cronjob

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/robfig/cron/v3"
)

func StartCleaner(db *sqlx.DB) {
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	c := cron.New(cron.WithLocation(location))

	c.AddFunc("0 0 * * *", func() {
		CleanAllTables(db)
	})

	c.Start()
}
