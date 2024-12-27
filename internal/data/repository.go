package data

import "github.com/sensfo/server/internal/encryption"

type SensfoDataSource struct {
	entity Repository
}

func (r *SensfoDataSource) Entity() Repository {
	return r.entity
}

func NewDataSource(encryption encryption.Encryption) DataSource {
	return &SensfoDataSource{entity: NewEntityData(encryption)}
}
