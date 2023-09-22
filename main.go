package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/sandy1206/GoKeyDB/config"
	"github.com/sandy1206/GoKeyDB/db"
	"github.com/sandy1206/GoKeyDB/web"
)

var (
	dbLocation = flag.String("db-location", "", "The Path to the bolt DB Database")
	httpAddr   = flag.String("http-addr", "127.0.0.1:3001", "The address to listen on for HTTP requests")
	configFile = flag.String("config-file", "sharding.toml", "Configuration file for sharding")
	shard      = flag.String("shard", "", "The name of the shard for data")
)

func parseFlags() {
	flag.Parse()

	if *dbLocation == "" {
		log.Fatal("db-location is required")
	}
	if *shard == "" {
		log.Fatalf("Must provide shard")
	}
}

func main() {
	parseFlags()

	var c config.Config
	if _, err := toml.DecodeFile(*configFile, &c); err != nil {
		log.Fatalf("Error decoding config file: %q", err)
	}

	var shardCount int
	var shardIdx int = -1
	var addrs = make(map[int]string)

	for _, s := range c.Shards {
		addrs[s.Idx] = s.Address

		if s.Idx+1 > shardCount {
			shardCount = s.Idx + 1
		}
		if s.Name == *shard {
			shardIdx = s.Idx
		}

	}

	if shardIdx < 0 {
		log.Fatalf("Shard %q not found in config file", *shard)
	}

	log.Printf("Shard Count is %d, current shard : %d", shardCount, shardIdx)

	db, close, err := db.NewDatabase(*dbLocation)
	if err != nil {
		log.Fatalf("Error opening database: %v: %q", *dbLocation, err)
	}
	defer close()

	srv := web.NewServer(db, shardCount, shardIdx, addrs)

	http.HandleFunc("/get", srv.GetHandler)
	http.HandleFunc("/set", srv.SetHandler)

	// hash := crc32.ChecksumIEEE([]byte("hello")) % uint32(shardCount)

	// Start the HTTP server
	http.ListenAndServe(*httpAddr, nil)

}
