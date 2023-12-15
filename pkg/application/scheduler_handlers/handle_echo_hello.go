package scheduler_handlers

import (
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/rs/zerolog/log"
)

func (svc *InAppScheduledJobService) ScheduleEchoHello() {
	// add a job to the scheduler
	_, err := svc.scheduler.NewJob(
		gocron.DurationJob(
			10*time.Second,
		),
		gocron.NewTask(
			func(text string, numb int64) {
				log.Info().Msgf("Now(): %d, StartedAt: %d, Message: %s", numb, time.Now().UnixMilli(), text)
			},
			"hi scheduler",
			time.Now().UnixMilli(),
		),
	)
	if err != nil {
		log.Err(err).Msg("app.scheduler.handle_echo_hello Error adding job")
	}
}
