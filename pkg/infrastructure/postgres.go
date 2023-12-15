package infrastructure

import (
	"github.com/rs/zerolog/log"

	// db
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(connStr string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("🛑 app.postgres.init: failed to connect postgres server")
	}

	log.Info().Msg("🔥 app.postgres.init: postgres connection established")

	// return postgres connection
	return db
}
