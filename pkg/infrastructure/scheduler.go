package infrastructure

import (
	"github.com/go-co-op/gocron/v2"
	"github.com/rs/zerolog/log"
)

func NewInAppScheduler() gocron.Scheduler {
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal().Err(err).Msg("app.scheduler.init: failed to initialize scheduler")
	}

	return s
}
