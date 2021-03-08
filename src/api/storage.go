package api

import (
	"catalog/src/db/postgres"
)

type Storage struct {
	db *postgres.DBConnections
}

func NewStorage(db *postgres.DBConnections) *Storage {
	return &Storage{db: db}
}
