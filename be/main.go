package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/syougo1209/b-match-server/config"
	"github.com/syougo1209/b-match-server/infrastructure/database"

	"github.com/labstack/echo/v4"
	router "github.com/syougo1209/b-match-server/interface"
)

func main() {

	ctx := context.Background()
	cfg, err := config.New()
	if err != nil {
		log.Printf("failed to new config: %v", err)
	}

	xdb, cleanup, err := database.NewDB(ctx, cfg)

	e, err := router.NewRouter(ctx, cfg, xdb)
	defer cleanup()
	if err != nil {
		log.Printf("failed to start database: %v", err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Helo, Worl! api")
	})

	go func() {
		addr := fmt.Sprintf(":%d", cfg.Port)
		if err := e.Start(addr); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
