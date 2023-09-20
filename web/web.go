package web

import (
	"fmt"
	"net/http"

	"github.com/sandy1206/GoKeyDB/db"
)

// Server is the main struct for the web server
type Server struct {
	db *db.Database
}

// NewServer creates a new server
func NewServer(db *db.Database) *Server {
	return &Server{db: db}
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

	err := s.db.SetKey(key, []byte(value))
	fmt.Fprintf(w, "%v", err)

}
