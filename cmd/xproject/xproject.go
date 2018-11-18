package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/radisvaliullin/cloudbilling/internal/handlers"

	"github.com/buaazp/fasthttprouter"
	"github.com/seatgeek/graceful_listener"
	"github.com/valyala/fasthttp"
)

func main() {
	router := fasthttprouter.New()
	apiPrefix := "/api/v1"

	router.GET(apiPrefix, handlers.GetSwaggerSpec)

	router.GET(apiPrefix+"/instance", handlers.GetAllInstances)

	router.GET(apiPrefix+"/instance/:id", handlers.GetInstancesByID)
	router.DELETE(apiPrefix+"/instance/:id", handlers.DeleteInstanceByID)
	router.GET(apiPrefix+"/instance/:id/history", handlers.GetInstanceHistoryByID)

	router.GET("/status", func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, "OK")
	})

	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatal("failed to open a socket: ", err)
	}

	gl := graceful_listener.NewGracefulListener(l, 5*time.Second)

	// Register system interrupt signals catching for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	go func() {
		log.Println("service has started")
		err := fasthttp.Serve(gl, router.Handler)
		if err != nil {
			log.Println(err)
			stop <- syscall.SIGINT
		}
	}()

	<-stop
	log.Println("shutting down the server...")
	gl.Close()
	log.Println("server gracefully stopped")
}
