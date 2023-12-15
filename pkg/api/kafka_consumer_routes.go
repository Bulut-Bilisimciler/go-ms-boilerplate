package api

import (
	"fmt"

	handlers "github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/application/kafka_handlers"
	"github.com/segmentio/kafka-go"
)

func InitKafkaConsumerRouter(svc *handlers.KafkaConsumerService) func(m *kafka.Message) error {
	// register consumer topic and keys
	consumeFn := func(m *kafka.Message) error {
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		// stringify incoming message
		msg := string(m.Value)
		key := string(m.Key)

		switch key {
		case "echo":
			return svc.ConsumeEchoIncomingText(msg)
		default:
			return svc.ConsumeEchoIncomingText(msg)
			// If you use Key based topic message seperation. You can use below code as default.
			// errors.New("no topic handlers found for topic " + m.Topic + "")
		}
	}

	return consumeFn
}
