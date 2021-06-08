package main

import (
	"github.com/techotron/online-quote-book/backend/api"
	"github.com/techotron/online-quote-book/backend/db"
	log "github.com/techotron/online-quote-book/backend/log"
)


func main() {
	router := api.Setup()

	err := db.SetupConn("api")
	if err != nil {
		log.Error("An error occured while setting up the db connection, check if db credentials are correct!")
		log.Error(err)
	}

	defer func() {
		err := db.CloseConn()
		if err != nil {
			log.Error(err)
		}
	}()

	err = db.MigrateUp()
	if err != nil {
		log.Error("An error occurred while running the database migration")
		log.Error(err)
	}

	startCron()

	err = router.Run(":5000")
	if err != nil {
		log.Error(err)
	}
}
