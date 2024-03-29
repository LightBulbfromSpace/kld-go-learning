package main

import (
	"github.com/LightBulbfromSpace/build_an_application"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {

	store, closeDB, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}

	defer closeDB()

	game := poker.NewTexasHoldem(store, poker.BlindAlerterFunc(poker.Alerter))
	server, err := poker.NewServer(store, game)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
