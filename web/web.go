package web

import (
	"fmt"
	"hash/fnv"
	"io"
	"net/http"

	"github.com/sandy1206/GoKeyDB/db"
)

// Server is the main struct for the web server
type Server struct {
	db         *db.Database
	shardCount int
	shardIdx   int
	addrs      map[int]string
}

func (s *Server) getShard(key string) int {
	h := fnv.New64()
	h.Write([]byte(key))
	return int(h.Sum64() % uint64(s.shardCount))
}

func (s *Server) redirect(shard int, w http.ResponseWriter, r *http.Request) {
	url := "http://" + s.addrs[shard] + r.RequestURI
	fmt.Fprintf(w, "redirecting from shard %d to shard %d (%q)\n", s.shardIdx, shard, url)

	resp, err := http.Get(url)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Error redirecting the request: %v", err)
		return
	}
	defer resp.Body.Close()

	io.Copy(w, resp.Body)
}

// NewServer creates a new server
func NewServer(db *db.Database, shardCount, shardIdx int, addrs map[int]string) *Server {
	return &Server{
		db:         db,
		shardCount: shardCount,
		shardIdx:   shardIdx,
		addrs:      addrs,
	}
}

func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")

	value, err := s.db.GetKey(key)
	shard := s.getShard(key)
	if shard != s.shardIdx {
		resp, err := http.Get("http://" + s.addrs[shard] + "/get?key=" + key)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Error getting value from shard %d: %v", shard, err)
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
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
