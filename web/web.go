package web

import "github.com/sandy1206/GoKeyDB/db"

type Server struct {
	db *db.Database
}

func NewServer(db *db.Database) *Server {
	return &Server{db: db}
}
