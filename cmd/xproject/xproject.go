package main

import (
	"log"

	"github.com/pavlov-tony/xproject/internal/handlers"

	"github.com/buaazp/fasthttprouter"
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

	log.Println("service is started")

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
