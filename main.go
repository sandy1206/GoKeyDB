package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/sandy1206/GoKeyDB/db"
	"github.com/sandy1206/GoKeyDB/web"
)

var (
	dbLocation = flag.String("db-location", "", "The Path to the bolt DB Database")
	httpAddr   = flag.String("http-addr", "127.0.0.1:3001", "The address to listen on for HTTP requests")
)

func parseFlags() {
	flag.Parse()

	if *dbLocation == "" {
		log.Fatal("db-location is required")
	}
}

func main() {
	parseFlags()
	db, close, err := db.NewDatabase(*dbLocation)
	if err != nil {
		log.Fatalf("Error opening database: %v: %q", *dbLocation, err)
	}
	defer close()

	srv := web.NewServer(db)

	http.HandleFunc("/get", srv.GetHandler)
	http.HandleFunc("/set", srv.SetHandler)

	http.ListenAndServe(*httpAddr, nil)

}
