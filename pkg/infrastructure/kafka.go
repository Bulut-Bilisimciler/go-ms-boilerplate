package infrastructure

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	// kafka
	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/internal/config"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func NewKafkaConsumerConnection(connStr string, cg string, topic string) *kafka.Reader {
	// Parse the URL
	parsed, err := url.Parse(connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("app.kafka.consumer.init: failed to parse kafka connection string")
	}

	// Extract components from the parsed URL
	username := parsed.User.Username()
	password, _ := parsed.User.Password()
	host := parsed.Hostname()
	port := parsed.Port()
	partitionStr := strings.TrimPrefix(parsed.Path, "/")

	// partition is digit string. cast to int.
	partition, err := strconv.Atoi(partitionStr)
	if err != nil {
		log.Fatal().Err(err).Msg("app.kafka.consumer.init: failed to parse kafka partition string")
	}

	// create kafka dialer
	dialer := &kafka.Dialer{
		Timeout:   time.Minute,
		DualStack: true,
		// SASLMechanism: mechanism,
	}

	// if env is "local" skip scram-auth.
	if config.C.App.Env != "local" {
		// scram.Mechanism
		mechanism, err := scram.Mechanism(scram.SHA512, username, password)
		if err != nil {
			log.Fatal().Err(err).Msg("app.kafka.consumer.init: failed to kafka scram auth")
		}

		dialer.SASLMechanism = mechanism
	}

	// create reader (consumer)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{
			host + ":" + port,
		},
		// kafka generic opts.
		GroupID:   cg,
		Topic:     topic,
		Dialer:    dialer,
		Partition: partition,
		// no wait for message bucket to be full.
		MinBytes: 1,   // 1 byte
		MaxBytes: 1e6, // 1MB
	})

	// log success
	log.Info().Msg("app.kafka.consumer.init: kafka consumer connected to topic " + topic + " connection initialized successfully.")

	return r
}

func NewKafkaProducerConnection(connStr string, defaultTopic string) *kafka.Writer {
	// Parse the URL
	parsed, err := url.Parse(connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("app.kafka.producer.init: failed to parse kafka connection string")
	}

	// Extract components from the parsed URL
	username := parsed.User.Username()
	password, _ := parsed.User.Password()
	host := parsed.Hostname()
	port := parsed.Port()

	// create kafka dialer
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		// SASLMechanism: mechanism,
	}

	// if env is "local" skip scram-auth.
	if config.C.App.Env != "local" {
		// scram.Mechanism
		mechanism, err := scram.Mechanism(scram.SHA512, username, password)
		if err != nil {
			// log.Fatal("INIT: kafka connection scram.Mechanism is failed: ", err)
			log.Fatal().Err(err).Msg("app.kafka.producer.init: failed to kafka scram auth")
		}

		dialer.SASLMechanism = mechanism
	}

	// create producer writer
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{
			host + ":" + port,
		},
		Topic:    defaultTopic,
		Balancer: &kafka.LeastBytes{},
		Dialer:   dialer,
		Async:    false,
	})

	// log success
	log.Info().Msg("app.kafka.producer.init: kafka producer with default topic " + defaultTopic + " connection initialized successfully.")

	return w
}
