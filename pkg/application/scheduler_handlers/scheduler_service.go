package scheduler_handlers

import (
	"context"

	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/infrastructure"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-co-op/gocron/v2"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

type InAppScheduledJobService struct {
	eventEmitter *infrastructure.EventEmitter
	natsClient   *nats.Conn
	kw           *kafka.Writer
	db           *gorm.DB
	rdb          *redis.Client
	s3Client     *s3.S3
	scheduler    gocron.Scheduler
}

func NewInAppScheduledJobService(
	eventEmitter *infrastructure.EventEmitter,
	kw *kafka.Writer,
	db *gorm.DB,
	rdb *redis.Client,
	s3Client *s3.S3,
	natsClient *nats.Conn,
	scheduler gocron.Scheduler,
) *InAppScheduledJobService {
	return &InAppScheduledJobService{
		eventEmitter: eventEmitter,
		db:           db,
		rdb:          rdb,
		s3Client:     s3Client,
		natsClient:   natsClient,
		kw:           kw,
		scheduler:    scheduler,
	}
}

func (svc *InAppScheduledJobService) InitAndRegisterJobs(ctx context.Context) {
	// all job assigned to global scheduler in gocron/v2
	// so its enough to just call Start() once
	svc.scheduler.Start()
}

func (svc *InAppScheduledJobService) ReleaseAllJobs() error {
	// drain and close
	err := svc.scheduler.Shutdown()
	return err
}
