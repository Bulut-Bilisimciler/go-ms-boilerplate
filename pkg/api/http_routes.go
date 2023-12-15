package api

import (
	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/internal/config"
	handlers "github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/application/http_handlers"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type RespondJson struct {
	Status  bool        `json:"status"`
	Intent  string      `json:"intent"`
	Message interface{} `json:"message"`
}

func respondJson(ctx *gin.Context, code int, intent string, message interface{}, err error) {
	if err == nil {
		ctx.JSON(code, RespondJson{
			Status:  true,
			Intent:  intent,
			Message: message,
		})
	} else {
		ctx.JSON(code, RespondJson{
			Status:  false,
			Intent:  intent,
			Message: err.Error(),
		})
	}
}

func InitRouter(svc *handlers.HTTPHandlerService, r *gin.Engine) *gin.Engine {
	// -- boilerplate routes (group)
	v1 := r.Group(config.API_PREFIX)

	// Utility routes "/ping" and "/metrics"
	// ping
	r.GET("/ping", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "pong"}) })
	// metrics
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// SECTION ROUTES
	// get user information from jwt
	v1.GET("/me", func(ctx *gin.Context) {
		code, data, err := svc.HandleGetUserInformation(ctx)
		respondJson(ctx, code, config.RN_PREFIX+".me", data, err)
	})

	return r
}
