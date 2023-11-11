package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8082"

type Config struct{}

func main() {
	app := Config{}
	log.Printf("Starting brocket service on port %s", webPort)

	//define http server

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}

}