package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yesilyurtburak/go-web-basics-3/pkg/config"
	"github.com/yesilyurtburak/go-web-basics-3/pkg/handlers"
)

const portNumber = "8080"
const ipAddress = "127.0.0.1"

var url = fmt.Sprintf("%s:%s", ipAddress, portNumber)

func main() {
	var app *config.AppConfig     // defines a new configuration variable `app`
	repo := handlers.NewRepo(app) // creates a new repo
	handlers.NewHandlers(repo)    // this assign the value `repo` to `Repo` variable inside the handlers.go

	// create and configure a new server
	srv := &http.Server{
		Addr:    url,
		Handler: routes(app),
	}

	// listen to the traffic for incoming http requests.
	fmt.Printf("Listening traffic at %s\n", url)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
