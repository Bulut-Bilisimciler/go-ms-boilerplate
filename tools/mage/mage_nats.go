//go:build mage
// +build mage

package main

import (
	"fmt"
	"time"

	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/internal/config"
	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/infrastructure"
	"github.com/magefile/mage/mg"
)

type Nats mg.Namespace

// Request makes a nats request to a subject Usage "mage nats:request SUBJECT MESSAGE"
func (Nats) Request(subject string, message string) error {
	ReadMageConfig()
	fmt.Printf("subject: %s, message: %s\n", subject, message)

	nc := infrastructure.NewNatsConnection(config.C.Nats.Url)
	defer nc.Close()

	subject = config.FN_PREFIX + "." + subject
	msg, err := nc.Request(subject, []byte(message), 1000*time.Millisecond)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("message received: ", string(msg.Data))

	return err
}

// RequestMany makes many nats request to a subject Usage "mage nats:requestMany AMOUNT SUBJECT MESSAGE"
func (Nats) RequestMany(amount int, subject string, message string) error {
	ReadMageConfig()
	fmt.Printf("amount: %d, subject: %s, message: %s\n", amount, subject, message)

	nc := infrastructure.NewNatsConnection(config.C.Nats.Url)
	defer nc.Close()

	subject = config.FN_PREFIX + "." + subject
	var err error
	for i := 0; i < amount; i++ {
		msg, err := nc.Request(subject, []byte(message), 1000*time.Millisecond)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%d message received: %s\n", i, string(msg.Data))
	}

	return err
}
