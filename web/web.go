package web

import (
	"fmt"
	"hash/fnv"
	"net/http"

	"github.com/sandy1206/GoKeyDB/db"
)

// Server is the main struct for the web server
type Server struct {
	db         *db.Database
	shardCount int
	shardIdx   int
}

// NewServer creates a new server
func NewServer(db *db.Database, shardCount, shardIdx int) *Server {
	return &Server{
		db:         db,
		shardCount: shardCount,
		shardIdx:   shardIdx,
	}
}

func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")

	value, err := s.db.GetKey(key)
	fmt.Fprintf(w, string(value), err)

}

func (s *Server) SetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	value := r.Form.Get("value")

	h := fnv.New64()
	h.Write([]byte(key))
	shardIdx := int(h.Sum64() % uint64(s.shardCount))

	err := s.db.SetKey(key, []byte(value))
	fmt.Fprintf(w, "%v", err)

}
