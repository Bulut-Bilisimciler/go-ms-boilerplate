package api

import (
	"errors"
	"fmt"

	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/internal/config"
	handlers "github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/application/nats_handlers"
	"github.com/nats-io/nats.go"
)

func InitNatsResponderRouter(svc *handlers.NatsReqReplyService) func(m *nats.Msg) (string, error) {
	// register consumer topic and keys
	respondFn := func(m *nats.Msg) (string, error) {
		fmt.Printf("NEW_MESSAGE /sub/message : /%v/%s\n", m.Subject, string(m.Data))
		// trim prefix of subject "FN_PREFIX."
		trimmed := m.Subject[len(config.FN_PREFIX)+1:]

		switch trimmed {
		case "echo":
			return svc.RespondEchoIncomingText(m)
		default:
			return "no subject handlers found", errors.New("no subject handlers found")
		}
	}

	// return the responder function
	return respondFn
}
