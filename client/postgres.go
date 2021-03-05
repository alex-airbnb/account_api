package client

import (
	"os"

	"github.com/alex-airbnb/account_api/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres Postgres DB Instance.
var Postgres *gorm.DB

// SetUpPostgres Set up the Postgres DB Instance.
func SetUpPostgres() error {
	Postgres, err := gorm.Open(
		postgres.Open(
			os.Getenv("DB_URL"),
		),
		&gorm.Config{},
	)

	if err != nil {
		return err
	}

	migrate(Postgres)

	return nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&entity.Account{})
}
