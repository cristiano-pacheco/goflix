package httpserver

import (
	"context"

	_ "github.com/cristiano-pacheco/goflix/docs"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/config"
	"github.com/cristiano-pacheco/goflix/pkg/httpserver"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/fx"
)

type HTTPServer struct {
	server *httpserver.HTTPServer
}

func NewHTTPServer(
	lc fx.Lifecycle,
	conf config.Config,
) *HTTPServer {
	corsConfig := httpserver.CorsConfig{
		AllowedOrigins:   conf.CORS.GetAllowedOrigins(),
		AllowedMethods:   conf.CORS.GetAllowedMethods(),
		AllowedHeaders:   conf.CORS.GetAllowedHeaders(),
		ExposedHeaders:   conf.CORS.GetExposedHeaders(),
		AllowCredentials: conf.CORS.AllowCredentials,
		MaxAge:           conf.CORS.MaxAge,
	}

	isOtelEnabled := true
	server := httpserver.NewHTTPServer(corsConfig, conf.App.Name, isOtelEnabled, conf.HTTPPort)

	httpServer := &HTTPServer{
		server: server,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			server.Run()
			return nil
		},
		OnStop: server.Shutdown,
	})

	return httpServer
}

func (s *HTTPServer) Router() *httprouter.Router {
	return s.server.Router()
}
