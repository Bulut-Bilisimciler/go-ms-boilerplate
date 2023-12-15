package api

import (
	"context"

	"go.uber.org/fx"

	// logger
	"github.com/rs/zerolog/log"

	handlers "github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/application/scheduler_handlers"
)

type InAppScheduledJobsState struct {
	IsStarted bool
	err       error
}

func NewInAppJobsStarter(lc fx.Lifecycle, svc *handlers.InAppScheduledJobService) *InAppScheduledJobsState {
	state := &InAppScheduledJobsState{
		IsStarted: false,
		err:       nil,
	}

	// init event handlers
	svc.InitEventHandlers()

	// init routes
	InitScheduledJobsRouter(svc)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			svc.InitAndRegisterJobs(ctx)
			log.Info().Msg("✅ [SCHEDULER] Server is started")
			return nil

		},
		OnStop: func(ctx context.Context) error {
			// stop the consumer
			err := svc.ReleaseAllJobs()
			if err != nil {
				log.Err(err).Msg("app.scheduler.release Error releasing all jobs")
				return err
			}

			log.Info().Msg("🛑 [SCHEDULER] Server is stopped")
			return nil
		},
	})

	return state
}
