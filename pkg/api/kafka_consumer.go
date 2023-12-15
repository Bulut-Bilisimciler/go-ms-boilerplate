package api

import (
	"context"

	"go.uber.org/fx"

	// logger
	"github.com/rs/zerolog/log"

	handlers "github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/application/kafka_handlers"
)

type KafkaConsumerListenerState struct {
	IsStarted bool
	err       error
}

func NewKafkaConsumerServer(lc fx.Lifecycle, svc *handlers.KafkaConsumerService) *KafkaConsumerListenerState {
	state := &KafkaConsumerListenerState{
		IsStarted: false,
		err:       nil,
	}

	// init event handlers
	svc.InitEventHandlers()

	consumeFn := InitKafkaConsumerRouter(svc)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go svc.InitAndRegisterConsumer(ctx, consumeFn)
			log.Info().Msg("✅ [KAFKA] CONSUMER Server is started")
			return nil

		},
		OnStop: func(ctx context.Context) error {
			// stop the consumer
			err := svc.CloseConsumer()
			if err != nil {
				log.Err(err).Msg("app.kafka.close Error closing kafka reader")
				return err
			}

			log.Info().Msg("🛑 [KAFKA] CONSUMER Server is stopped")
			return nil
		},
	})

	return state
}
