package server

import (
	"context"
	"fmt"
	"net/http"
	"person-predicator/internal/logger"
	"person-predicator/internal/server/handlers/persons"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type HTTP struct {
	cfg           *Config
	HTTPServer    *http.Server
	personHandler *persons.PersonHandler
}

func NewHTTP(cfg *Config, ph *persons.PersonHandler) *HTTP {
	return &HTTP{cfg: cfg, personHandler: ph}
}

// @title		Device Manager API
// @version	1.0
// @description
// @BasePath
func (s *HTTP) StartHTTP(ctx context.Context) {
	r := gin.Default()
	// docs.SwaggerInfo.BasePath = "/"
	// docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)

	{
		r.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Test response")
		})

		r.POST("/person", s.personHandler.Add)
		r.GET("/device/:uuid", s.personHandler.Get)

		// r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	s.HTTPServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port),
		Handler: r,
	}

	logger.Logger.Info("Starting HTTP server ...")
	err := s.HTTPServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Logger.Error("Failed to start HTTP server", zap.Error(err))
	}

	// shutdownCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	// defer cancel()
	// logger.Logger.Info("Shutting down HTTP server ...")
	// if err := s.HTTPServer.Shutdown(shutdownCtx); err != nil {
	// 	logger.Logger.Error("Failed to shutdown HTTP server", zap.Error(err))
	// }
}
