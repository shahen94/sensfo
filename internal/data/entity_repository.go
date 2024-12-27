package data

import (
	"errors"

	"github.com/sensfo/server/internal/encryption"
)

// ─────────────────────────────────────────────────────────────────────────────

type EntityRepository struct {
	encryption encryption.Encryption
	entities   map[string]string
}

func (e *EntityRepository) Store(key string, value string) (encryption.Result, error) {
	result, err := e.encryption.Encrypt(value)

	if err != nil {
		return nil, err
	}

	e.entities[result.Content()] = key

	return result, nil
}

func (e *EntityRepository) Retrieve(key string) (string, error) {
	bias := e.encryption.ComputeBias(key)

	stored, err := e.encryption.Decrypt(key, bias)

	if err != nil {
		return "", err
	}

	value, ok := e.entities[stored]

	if !ok {
		return "", errors.New("entity not found")
	}

	return value, nil
}

// ─────────────────────────────────────────────────────────────────────────────

func NewEntityData(encryption encryption.Encryption) Repository {
	return &EntityRepository{
		encryption: encryption,
		entities:   make(map[string]string),
	}
}
