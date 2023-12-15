package http_handlers

import (
	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/infrastructure"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

type HTTPHandlerService struct {
	eventEmitter *infrastructure.EventEmitter
	db           *gorm.DB
	rdb          *redis.Client
	s3Client     *s3.S3
	natsClient   *nats.Conn
	kw           *kafka.Writer
}

func NewHTTPHandlerService(
	eventEmitter *infrastructure.EventEmitter,
	db *gorm.DB,
	rdb *redis.Client,
	s3Client *s3.S3,
	natsClient *nats.Conn,
	kw *kafka.Writer,
) *HTTPHandlerService {
	return &HTTPHandlerService{
		eventEmitter: eventEmitter,
		db:           db,
		rdb:          rdb,
		s3Client:     s3Client,
		natsClient:   natsClient,
		kw:           kw,
	}
}
