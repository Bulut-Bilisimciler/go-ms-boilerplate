package kafka_handlers

import (
	"fmt"
)

func (svc *KafkaConsumerService) ConsumeEchoIncomingText(message string) error {
	// dont try to marshall or unmarshall string. Just echo
	// e.g. svc.rdb.Set(c.Request.Context(), "test", message, 0) ...
	fmt.Printf("ConsumeEchoIncomingTestText ECHO: %s\n", message)

	// return nil
	return nil
}

// example usage of redis: try to get redis keys *
// keys, err := svc.rdb.Keys(c.Request.Context(), "*").Result()
// if err != nil {
// 	return http.StatusInternalServerError, nil, err
// }
