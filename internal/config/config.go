package config

import (
	"github.com/rs/zerolog/log"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

const (
	// server name
	APP_NAME string = "bb.app.boilerplate"
	// server description
	APP_DESCRIPTION string = "bb.app.boilerplate is a microservice boilerplate for Bulut-Bilisimciler"
	// API Route Prefix if exist (e.g "" or "/api-boilerplate")
	API_PREFIX string = "/api-boilerplate"
	// Resource Name Prefix if exist (e.g "bb.app.boilerplate")
	RN_PREFIX string = "bb.app.boilerplatesvc"
	// Nats FN Prefix if exist (e.g "bb.app.boilerplate")
	FN_PREFIX string = "bb.fn.boilerplatesvc"
)

type config struct {
	App struct {
		Env     string `mapstructure:"env"`
		Port    string `mapstructure:"port"`
		Version string `mapstructure:"version"`
	}

	Auth struct {
		JwtPub string `mapstructure:"jwt_pub"`
	}

	Recaptcha struct {
		Host   string `mapstructure:"host"`
		Secret string `mapstructure:"secret"`
	}

	Db struct {
		Url string `mapstructure:"url"`
	}

	Cache struct {
		Url string `mapstructure:"url"`
	}

	Broker struct {
		Url                      string `mapstructure:"url"`
		ConsumerGroup            string `mapstructure:"consumer_group"`
		ConsumerTopicPrefix      string `mapstructure:"consumer_topic_prefix"`
		TopicToConsume           string `mapstructure:"topic_to_consume"`
		ProducerDeadLetterPrefix string `mapstructure:"producer_dead_letter_prefix"`
	} `mapstructure:"broker"`

	Nats struct {
		Url            string `mapstructure:"url"`
		ResponderGroup string `mapstructure:"responder_group"`
	} `mapstructure:"nats"`

	Cdn struct {
		Url    string `mapstructure:"url"`
		Bucket string `mapstructure:"bucket"`
	} `mapstructure:"cdn"`

	PromProbe struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"promprobe"`
}

var C config
var Pwd string
var IsProd bool

func ReadConfig(absoluteConfigDir string, echoEnv bool) {
	Config := &C
	viper.SetConfigName(".env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(absoluteConfigDir)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("app.config.load error cannot read .env.yaml file")
	}

	if err := viper.Unmarshal(Config); err != nil {
		log.Fatal().Err(err).Msg("app.config.load error cannot unmarshal .env.yaml file")
	}

	// checking env...
	IsProd = C.App.Env == "prod" || C.App.Env == "production"

	if echoEnv {
		// print config
		spew.Dump(C)
	}
}
