//go:build mage
// +build mage

package main

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"strconv"

	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/internal/config"
	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/infrastructure"
	"github.com/magefile/mage/mg"
	"github.com/segmentio/kafka-go"
	"github.com/sourcegraph/conc"
)

type Kafka mg.Namespace

// CreateTopic creates a topic Usage "mage kafka:createTopic TOPIC"
func (Kafka) CreateTopic(topic string) error {
	ReadMageConfig()
	fmt.Printf("topic: %s\n", topic)

	// topic should be only alphanumeric and "." and "-" characters
	url, _ := url.Parse(config.C.Broker.Url)
	conn, err := kafka.Dial("tcp", net.JoinHostPort(url.Hostname(), url.Port()))
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("topic created")
	return err
}

// CreateTopic creates a topic Usage "mage kafka:publish TOPIC MESSAGE"
func (Kafka) Publish(topic string, message string) error {
	ReadMageConfig()
	fmt.Printf("topic: %s, message: %s\n", topic, message)

	kw := infrastructure.NewKafkaProducerConnection(config.C.Broker.Url, topic)
	defer kw.Close()

	err := kw.WriteMessages(context.Background(), kafka.Message{
		Value: []byte(message),
	})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("message written")

	return err
}

// Consume consume messages from a topic Usage "mage kafka:consume TOPIC"
func (Kafka) Consume(cg string, topic string) error {
	ReadMageConfig()
	fmt.Printf("consumergroup: %s, topic: %s\n", cg, topic)

	kw := infrastructure.NewKafkaConsumerConnection(config.C.Broker.Url, cg, topic)
	defer kw.Close()

	var wg conc.WaitGroup
	defer wg.Wait()

	wg.Go(func() {
		for {
			// until it recieves a interrupt signal
			m, err := kw.ReadMessage(context.Background())
			if err != nil {
				panic(err.Error())
			}
			fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		}

	})

	return nil
}
