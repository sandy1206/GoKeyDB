package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/sandy1206/GoKeyDB/db"
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

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		key := r.Form.Get("key")
		value, err := db.GetKey(key)
		fmt.Println(value, err)

	})

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		key := r.Form.Get("key")
		value := r.Form.Get("value")
		err := db.SetKey(key, []byte(value))
		fmt.Println(err)

	})

	http.ListenAndServe(*httpAddr, nil)

}
