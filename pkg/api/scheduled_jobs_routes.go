package api

import (
	handlers "github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/application/scheduler_handlers"
)

func InitScheduledJobsRouter(svc *handlers.InAppScheduledJobService) {

	// Job Echo: echo hello from scheduler every 5 seconds
	svc.ScheduleEchoHello()

	// TODO: add more jobs here

}
