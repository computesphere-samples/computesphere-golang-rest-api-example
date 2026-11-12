package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"computesphere.com/computesphere-golang-rest-api-example/pkg/logging"
	"computesphere.com/computesphere-golang-rest-api-example/routers"
)

func init() {
	logging.Setup()
}

func main() {
	// The platform sets PORT to the service's port; 8000 when it isn't set.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	endPoint := ":" + port
	server := &http.Server{
		Addr:         endPoint,
		Handler:      routers.InitRouter(),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
