package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"

	server "github.com/dhimasan0206/adapter/sdk/server/go"
	"github.com/dhimasan0206/tracing"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type Services struct {
	InventoryService InventoryService
}

type API interface {
	Run(port string, services Services) error
}

type api struct{}

func Default() API {
	return &api{}
}

func (a *api) Run(port string, services Services) error {
	log.Printf("Server started")

	// Tracing
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	otelShutdown, err := tracing.SetupOTelSDK(ctx)
	if err != nil {
		return err
	}

	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	router := gin.Default()

	router.Use(otelgin.Middleware(viper.GetString("APP_NAME")))

	routes := NewApiHandleFunctions(services)

	router = server.NewRouterWithGinEngine(router, routes)

	return router.Run(fmt.Sprintf(":%v", port))
}
