package infrastructure

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	// cache
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func NewRedisCacheConnection(appCtx context.Context, connStr string) *redis.Client {
	// Parse the URL
	parsed, err := url.Parse(connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("app.redis.init: failed to parse redis connection string")
	}

	// Extract components from the parsed URL
	username := parsed.User.Username()
	password, _ := parsed.User.Password()
	host := parsed.Hostname()
	port := parsed.Port()
	index := strings.TrimPrefix(parsed.Path, "/")

	// convert db to int
	namespace, err := strconv.Atoi(index)
	if err != nil {
		log.Fatal().Err(err).Msg("app.redis.init: failed to parse redis db to int")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Username: username,
		Password: password,
		DB:       namespace,
	})

	// ping redis for check connection
	_, err = rdb.Ping(appCtx).Result()
	if err != nil {
		log.Fatal().Err(err).Msg("app.redis.init: failed to ping redis server")
	}

	// return redis connection
	return rdb
}
