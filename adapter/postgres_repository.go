package adapter

import "gorm.io/gorm"

type postgres struct {
	client *gorm.DB
}

// NewPostgres Create a new instance of Postgres Adapter.
func NewPostgres(db *gorm.DB) RepositoryPort {
	return &postgres{
		client: db,
	}
}

// Create Create a new recod in the DB.
func (p *postgres) Create(model interface{}) error {
	return p.client.Create(model).Error
}
