package nats_handlers

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func (svc *NatsReqReplyService) RespondEchoIncomingText(m *nats.Msg) (string, error) {
	// dont try to marshall or unmarshall string. Just echo
	// e.g. svc.rdb.Set(c.Request.Context(), "test", message, 0) ...
	fmt.Printf("RespondEchoIncomingTestText ECHO: %s\n", string(m.Data))

	// return nil
	return string(m.Data), nil
}

// example usage of redis: try to get redis keys *
// keys, err := svc.rdb.Keys(c.Request.Context(), "*").Result()
// if err != nil {
// 	return http.StatusInternalServerError, nil, err
// }
