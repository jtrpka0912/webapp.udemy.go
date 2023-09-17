package main

import (
	"encoding/gob"
	"flag"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"webapp.udemy.go/pkg/data"
	"webapp.udemy.go/pkg/db"
)

type application struct {
	DSN     string
	DB      db.PostgresConn
	Session *scs.SessionManager
}

func main() {
	gob.Register(data.User{})
	// Set up an app config
	app := application{}

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5", "Postgres Connection")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.DB = db.PostgresConn{DB: conn}

	// Get a session manager
	app.Session = getSession()

	// Print out a message
	log.Println("Starting server on port 8080")

	// Start the server
	err = http.ListenAndServe(":8080", app.routes())

	if err != nil {
		log.Fatal(err)
	}
}
