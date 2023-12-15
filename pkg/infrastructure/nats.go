package infrastructure

import (
	"net/url"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

func NewNatsConnection(connStr string) *nats.Conn {
	// Parse the URL
	parsed, err := url.Parse(connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("app.nats.init: failed to parse nats connection string")
	}

	// Extract components from the parsed URL
	scheme := parsed.Scheme // e.g. "s3"
	username := parsed.User.Username()
	password, _ := parsed.User.Password()
	host := parsed.Hostname()
	port := parsed.Port()

	nc, err := nats.Connect(scheme+"://"+host+":"+port, nats.UserInfo(username, password))
	if err != nil {
		log.Fatal().Err(err).Msg("app.nats.init: failed to connect nats server")
	}

	// return nats connection
	return nc
}
