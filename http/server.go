package http

import (
	"context"
	"core/infra/config"
	"core/infra/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Start() {
	g := gin.Default()

	//if err := middlewares.Attach(e); err != nil {
	//	logger.Error("error occur when attaching middlewares", err)
	//	os.Exit(1)
	//}

	Init(g.Group("api"))
	port := config.App().Port

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: g,
	}
	/// start http server
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()
	fmt.Println("server listening on port : ", port)

	// graceful shutdown
	GracefulShutdown(srv)
}

// server will gracefully shutdown within 5 sec
func GracefulShutdown(srv *http.Server) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	logger.Info("server shutdowns gracefully")
}
