package main

import (
	"log"
	"net/http"
	"time"

	"computesphere.com/computesphere-golang-rest-api-example/pkg/logging"
	"computesphere.com/computesphere-golang-rest-api-example/routers"
)

func init() {
	logging.Setup()
}

func main() {
	endPoint := ":8000"
	server := &http.Server{
		Addr:         endPoint,
		Handler:      routers.InitRouter(),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}

// prod preview e2e validation

// re-trigger prod preview

// prod preview after billing active

// preview after base manifest exists

// re-trigger for route (port now in DB)
