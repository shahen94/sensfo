package data

import "github.com/sensfo/server/internal/encryption"

type Repository interface {
	Store(string, string) (encryption.Result, error)
	Retrieve(string) (string, error)
}

type DataSource interface {
	Entity() Repository
}
