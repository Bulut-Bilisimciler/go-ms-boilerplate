package kafka_handlers

import (
	"context"

	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/infrastructure"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/sourcegraph/conc"
	"gorm.io/gorm"
)

type KafkaConsumerService struct {
	eventEmitter *infrastructure.EventEmitter
	kr           *kafka.Reader
	db           *gorm.DB
	rdb          *redis.Client
	s3Client     *s3.S3
	natsClient   *nats.Conn
}

func NewKafkaConsumerService(
	eventEmitter *infrastructure.EventEmitter,
	kr *kafka.Reader,
	db *gorm.DB,
	rdb *redis.Client,
	s3Client *s3.S3,
	natsClient *nats.Conn,
) *KafkaConsumerService {
	return &KafkaConsumerService{
		eventEmitter: eventEmitter,
		db:           db,
		rdb:          rdb,
		s3Client:     s3Client,
		natsClient:   natsClient,
		kr:           kr,
	}
}

func (svc *KafkaConsumerService) InitAndRegisterConsumer(ctx context.Context, consumeFn func(m *kafka.Message) error) {
	var wg conc.WaitGroup
	defer wg.Wait()
	kafkaContext := context.Background()

	wg.Go(func() {
		for {
			// read kafka message
			m, err := svc.kr.ReadMessage(kafkaContext)
			if err != nil {
				log.Err(err).Msg("app.kafka.read error reading kafka message")
				break
			}

			// process kafka message
			if err := consumeFn(&m); err != nil {
				log.Err(err).Msg("app.kafka.consume error processing kafka message")

				// emit error event
				// TODO: beautify error message
				svc.eventEmitter.EmitSync(kafkaContext, "kafka.message.error", err)
				continue
			}

			// emit success event
			// TODO: beautify success message and data
			svc.eventEmitter.EmitAsync(kafkaContext, "kafka.message.received", m)

		}
	})
}

func (k *KafkaConsumerService) CloseConsumer() error {
	err := k.kr.Close()

	return err
}
