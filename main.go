package main

import (
	"flag"
	"log"
	"net/http"

	bolt "go.etcd.io/bbolt"
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
	db, err := bolt.Open(*dbLocation, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "key is required", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))

	})

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("Key:")
		if key == "" {
			http.Error(w, "key is required", http.StatusBadRequest)
			return
		}
	})

	http.ListenAndServe(*httpAddr, nil)

}
