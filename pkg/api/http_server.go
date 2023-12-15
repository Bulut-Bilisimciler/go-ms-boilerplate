package api

import (
	"context"
	"net"
	"net/http"

	// http
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	// logger
	"github.com/gin-contrib/logger"
	"github.com/rs/zerolog/log"

	// in-app
	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/internal/config"
	handlers "github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/application/http_handlers"

	// swagger
	swagdocs "github.com/Bulut-Bilisimciler/go-ms-boilerplate/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewGinHTTPServer(lc fx.Lifecycle, svc *handlers.HTTPHandlerService) *http.Server {
	port := config.C.App.Port

	// check env and set gin mode
	gin.SetMode(gin.DebugMode)
	if config.IsProd {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())

	// check env and set swagger
	if !config.IsProd {
		// enable gin logger
		router.Use(logger.SetLogger())

		// enable swagger
		swagdocs.SwaggerInfo.BasePath = config.API_PREFIX
		router.GET(config.API_PREFIX+"/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	// init event handlers
	svc.InitEventHandlers()

	// init routes
	InitRouter(svc, router)

	// create http server
	srv := &http.Server{Addr: ":" + port, Handler: router}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr) // the web server starts listening on 8080
			if err != nil {
				log.Info().Err(err).Msg("[HTTP] HTTP Server is not started due to error")
				return err
			}
			// process an incoming request in a go routine
			go srv.Serve(ln)
			log.Info().Msg("✅ [HTTP] HTTP Server is started")

			return nil

		},
		OnStop: func(ctx context.Context) error {
			// stop the web server
			srv.Shutdown(ctx)
			log.Info().Msg("🛑 [HTTP] HTTP Server is stopped")
			return nil
		},
	})

	return srv
}
