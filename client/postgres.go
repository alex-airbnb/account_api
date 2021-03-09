package client

import (
	"os"

	"github.com/alex-airbnb/account_api/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetUpPostgres Set up the Postgres DB Instance.
func SetUpPostgres() (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.Open(
			os.Getenv("DB_URL"),
		),
		&gorm.Config{},
	)

	if err != nil {
		return db, err
	}

	migrate(db)

	return db, err
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&entity.Account{})
}
