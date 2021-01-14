package main

import (
	"github.com/sirupsen/logrus"

	"github.com/BenJetson/netwatch/store"
)

func main() {
	log := logrus.New()
	dbLog := log.WithField("source", "database")

	db, err := store.NewDataStore(dbLog, "./store.db")
	if err != nil {
		log.Fatalf("could not start DB: %+v\n", err)
	}

	defer db.Close()
}
