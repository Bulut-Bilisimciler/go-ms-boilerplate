package nats_handlers

import (
	"context"

	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/internal/config"
	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/infrastructure"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

type NatsReqReplyService struct {
	eventEmitter *infrastructure.EventEmitter
	natsClient   *nats.Conn
	kw           *kafka.Writer
	db           *gorm.DB
	rdb          *redis.Client
	s3Client     *s3.S3
}

func NewNatsReqReplyService(
	eventEmitter *infrastructure.EventEmitter,
	kw *kafka.Writer,
	db *gorm.DB,
	rdb *redis.Client,
	s3Client *s3.S3,
	natsClient *nats.Conn,
) *NatsReqReplyService {
	return &NatsReqReplyService{
		eventEmitter: eventEmitter,
		db:           db,
		rdb:          rdb,
		s3Client:     s3Client,
		natsClient:   natsClient,
		kw:           kw,
	}
}

func (svc *NatsReqReplyService) InitAndRegisterResponder(ctx context.Context, respondeFn func(m *nats.Msg) (string, error)) {
	// Replies
	svc.natsClient.QueueSubscribe(config.FN_PREFIX+".*", config.C.Broker.ConsumerGroup, func(m *nats.Msg) {
		// process nats message
		response, err := respondeFn(m)
		if err != nil {
			log.Err(err).Msg("app.nats.respond error processing nats message")

			// emit error event
			// TODO: ...
		}

		// process response
		m.Respond([]byte(response))

		// emit success event
		// TODO: ...

	})
}

func (svc *NatsReqReplyService) CloseResponder() error {
	// drain and close
	err := svc.natsClient.Drain()
	return err
}
