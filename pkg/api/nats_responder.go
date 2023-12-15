package api

import (
	"context"

	"go.uber.org/fx"

	// logger
	"github.com/rs/zerolog/log"

	handlers "github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/application/nats_handlers"
)

type NatsResponderState struct {
	IsStarted bool
	err       error
}

func NewNatsResponderServer(lc fx.Lifecycle, svc *handlers.NatsReqReplyService) *NatsResponderState {
	state := &NatsResponderState{
		IsStarted: false,
		err:       nil,
	}

	// init event handlers
	svc.InitEventHandlers()

	natsRouterFns := InitNatsResponderRouter(svc)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			svc.InitAndRegisterResponder(ctx, natsRouterFns)
			log.Info().Msg("✅ [NATS] RESPONDER Server is started")
			return nil

		},
		OnStop: func(ctx context.Context) error {
			// stop the consumer
			err := svc.CloseResponder()
			if err != nil {
				log.Err(err).Msg("app.nats.close Error closing nats connection")
				return err
			}

			log.Info().Msg("🛑 [NATS] RESPONDER Server is stopped")
			return nil
		},
	})

	return state
}
