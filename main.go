package main

import (
	"flag"
	"fmt"
	"github.com/fitrah-firdaus/notes/db"
	"github.com/fitrah-firdaus/notes/notes/repository"
	"os"

	"github.com/fitrah-firdaus/notes/config"
	"github.com/fitrah-firdaus/notes/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()

	// init notes repository
	repository.NewMysqlNotesRepository(db.GetDB())

	server.Init()
}
