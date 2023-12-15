package kafka_handlers

import (
	"context"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_kafka_event_total",
		Help: "The total number of processed kafka events",
	})
	opsNotProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_error_processed_kafka_event_total",
		Help: "The total number of procession error kafka events",
	})
)

func (svc *KafkaConsumerService) InitEventHandlers() {
	svc.eventEmitter.On("kafka.message.received", func(ctx context.Context, payload interface{}) {
		fmt.Printf("👍🏻 MESSAGE_RECEIVED: %v\n", payload)
		opsProcessed.Inc()
	})
	svc.eventEmitter.On("kafka.message.error", func(ctx context.Context, payload interface{}) {
		fmt.Printf("👎🏻 MESSAGE_ERROR: %v\n", payload)
		opsNotProcessed.Inc()
	})

	// svc.eventEmitter.On("user.created", svc.UserCreatedEventHandler)
	// svc.eventEmitter.On("user.updated", svc.UserUpdatedEventHandler)
	// svc.eventEmitter.On("user.deleted", svc.UserDeletedEventHandler)
	// svc.eventEmitter.On("user.password.changed", svc.UserPasswordChangedEventHandler)
	// svc.eventEmitter.On("user.password.forgot", svc.UserPasswordForgotEventHandler)
	// svc.eventEmitter.On("user.password.reset", svc.UserPasswordResetEventHandler)
	// svc.eventEmitter.On("user.email.changed", svc.UserEmailChangedEventHandler)
	// svc.eventEmitter.On("user.email.verified", svc.UserEmailVerifiedEventHandler)
	// svc.eventEmitter.On("user.phone.changed", svc.UserPhoneChangedEventHandler)
	// svc.eventEmitter.On("user.phone.verified", svc.UserPhoneVerifiedEventHandler)
	// svc.eventEmitter.On("user.phone.forgot", svc.UserPhoneForgotEventHandler)
	// svc.eventEmitter.On("user.phone.reset", svc.UserPhoneResetEventHandler)
	// svc.eventEmitter.On("user.otp.changed", svc.UserOTPChangedEventHandler)
	// svc.eventEmitter.On("user.otp.verified", svc.UserOTPVerifiedEventHandler)
	// svc.eventEmitter.On("user.otp.forgot", svc.UserOTPForgotEventHandler)
	// svc.eventEmitter.On("user.otp.reset", svc.UserOTPResetEventHandler)
	// svc.eventEmitter.On("user.profile.updated", svc.UserProfileUpdatedEventHandler)
	// svc.eventEmitter.On("user.profile.deleted", svc.UserProfileDeletedEventHandler)
	// svc.eventEmitter.On("user.profile.created", svc.UserProfileCreatedEventHandler)
	// svc.eventEmitter.On("user.profile.verified", svc.UserProfileVerifiedEventHandler)
	// svc.eventEmitter.On("user.profile.forgot", svc.UserProfileForgotEventHandler)
	// svc.eventEmitter.On("user.profile.reset", svc.UserProfileResetEventHandler)
	// svc.eventEmitter.On("user.profile.password.changed", svc.UserProfilePasswordChangedEventHandler)
	// svc.eventEmitter.On("user.profile.password.forgot", svc.UserProfilePasswordForgotEventHandler)
	// svc.eventEmitter.On("user.profile.password.reset", svc.UserProfilePasswordResetEventHandler)
	// svc.eventEmitter.On("user.profile.email.changed", svc.UserProfileEmailChangedEventHandler)
	// svc.eventEmitter.On("user.profile.email.verified", svc.UserProfileEmailVerifiedEventHandler)
	// svc.eventEmitter.On("user.profile.phone.changed", svc.UserProfilePhoneChangedEventHandler)
	// svc.eventEmitter.On("user.profile.phone.verified", svc.UserProfilePhoneVerifiedEventHandler)
	// svc.eventEmitter.On("user.profile.phone.forgot", svc.UserProfilePhoneForgotEventHandler)
	// svc.eventEmitter.On("user.profile.phone.reset", svc.UserProfilePhoneResetEventHandler)
	// svc.eventEmitter.On("user.profile.otp.changed", svc.UserProfileOTPChangedEventHandler)
	// svc.eventEmitter.On("user.profile.otp.verified", svc.UserProfileOTPVerifiedEventHandler)
	// svc.eventEmitter.On("user.profile.otp.forgot", svc.UserProfileOTPForgotEventHandler)
	// svc.eventEmitter.On("user.profile.otp.reset", svc.UserProfileOTPResetEventHandler)
}
