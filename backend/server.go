package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

)

func main() {
	ctx := context.Background()

	defer func() {
		// Capturing panic of starting process.
		if x := recover(); x != nil {
			debug.PrintStack()
			errStr := fmt.Sprint(x)
			// recovering from a panic; x contains whatever was passed to panic()
			log.Errorf(ctx, fmt.Sprintf("run time panic: %s", errStr))

			// just want to log a panic, so panic again
			panic(x)
		}
	}()

	log.Infof(ctx, "Will start a process for %s", gin.Mode())

	engine, cleanFn, conf := initWithWire()
	defer func() {
		log.Infof(ctx, "Start cleanup of app...")
		cleanFn()
		log.Infof(ctx, "Done cleanup")
	}()

	startWithGracefullyShutdown(ctx, engine, conf.Server.Port)
}

func initWithWire() (*gin.Engine, func(), config.Config) {
	var app *gin.Engine
	var appCleanFn func()
	conf := di.InitConfig()

	app, appCleanFn = di.InitAppServerDependency(conf)
	return app, appCleanFn, conf
}

func startWithGracefullyShutdown(ctx context.Context, engine *gin.Engine, port string) {
	addr := fmt.Sprintf(":%s", port)

	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}

	log.Infof(ctx, "will launch on %s", addr)

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf(ctx, "listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Infof(ctx, "Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	shutdown := make(chan error)
	go func() {
		shutdown <- srv.Shutdown(ctx)
	}()

	select {
	case <-ctx.Done():
		log.Infof(ctx, "timeout of 10 seconds.")
	case err := <-shutdown:
		if err == nil {
			log.Infof(ctx, "Shutdown succeeded")
		} else {
			log.Fatalf(ctx, "Server Shutdown: %s", err)
		}
	}
}