package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/internal/config"
	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/api"
	handlerhttp "github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/application/http_handlers"
	handlerkafka "github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/application/kafka_handlers"
	handlernats "github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/application/nats_handlers"
	handlerscheduler "github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/application/scheduler_handlers"
	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/infrastructure"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	// open telemetry
)

// Path: bb-boilerplate-service
// @Title BB Boilerplate API
// @Description bb.app.boilerplate is a microservice boilerplate for Bulut-Bilisimciler
// @Version 1.0.0
// @Schemes http https
// @BasePath /api-boilerplate

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

const (
// for in-app global variables go "/internal/config/config.go"
)

func main() {
	// init app context
	appCtx := context.Background()

	// get config dir from flag or default "os.Getwd()"
	// e.g. "go run main.go --config-dir=/app/config"
	configDir := flag.String("config-dir", "", "absolute path of config directory like /app/config")
	flag.Parse()

	// init logger
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// get current working directory
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Msg("app.config.load error cannot get current working directory osgetwd")
	}
	config.Pwd = pwd

	// if configDir empty use default "$PWD/internal/config"
	absoluteDir := pwd + "/internal/config"
	if configDir != nil {
		if *configDir != "" {
			absoluteDir = *configDir
		}
	}

	// load and parse application envs
	config.ReadConfig(absoluteDir, false)
	if config.IsProd {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}

	// 3th party connections
	// e.g. event emitter, db, redis, kafka, nats, s3, etc.
	ee := infrastructure.NewEventEmitter()
	db := infrastructure.NewPostgresDB(config.C.Db.Url)
	rdb := infrastructure.NewRedisCacheConnection(appCtx, config.C.Cache.Url)
	s3Client := infrastructure.NewS3Session(config.C.Cdn.Url)
	natsClient := infrastructure.NewNatsConnection(config.C.Nats.Url)
	kw := infrastructure.NewKafkaProducerConnection(config.C.Broker.Url, config.C.Broker.ProducerDeadLetterPrefix+".ping")
	kr := infrastructure.NewKafkaConsumerConnection(config.C.Broker.Url, config.C.Broker.ConsumerGroup, config.C.Broker.ConsumerTopicPrefix+"."+config.C.Broker.TopicToConsume)
	s := infrastructure.NewInAppScheduler()

	// init http handlers
	svcHttp := handlerhttp.NewHTTPHandlerService(
		ee,
		db,
		rdb,
		s3Client,
		natsClient,
		kw,
	)

	// init kafka handlers
	svcKafka := handlerkafka.NewKafkaConsumerService(
		ee,
		kr,
		db,
		rdb,
		s3Client,
		natsClient,
	)

	// init nats handlers
	svcNats := handlernats.NewNatsReqReplyService(
		ee,
		kw,
		db,
		rdb,
		s3Client,
		natsClient,
	)

	// init in-app-scheduler handlers
	svcScheduler := handlerscheduler.NewInAppScheduledJobService(
		ee,
		kw,
		db,
		rdb,
		s3Client,
		natsClient,
		s,
	)

	fx.New(
		fx.Provide(
			func() *handlerhttp.HTTPHandlerService {
				log.Trace().Msg("🚀 STARTED HTTPHandlerService")
				return svcHttp
			},
			func() *handlerkafka.KafkaConsumerService {
				log.Trace().Msg("🚀 STARTED KafkaConsumerService")
				return svcKafka
			},
			func() *handlernats.NatsReqReplyService {
				log.Trace().Msg("🚀 STARTED NatsReqReplyService")
				return svcNats
			},
			func() *handlerscheduler.InAppScheduledJobService {
				log.Trace().Msg("🚀 STARTED InAppScheduledJobService")
				return svcScheduler
			},
			api.NewKafkaConsumerServer,
			api.NewNatsResponderServer,
			api.NewInAppJobsStarter,
			api.NewGinHTTPServer,
		),
		fx.Invoke(
			func(*api.KafkaConsumerListenerState) {},
			func(*api.NatsResponderState) {},
			func(*api.InAppScheduledJobsState) {},
			func(*http.Server) {},
		),
	).Run()
}
