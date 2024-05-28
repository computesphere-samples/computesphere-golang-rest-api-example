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
