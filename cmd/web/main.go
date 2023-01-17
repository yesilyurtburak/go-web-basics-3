package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yesilyurtburak/go-web-basics-3/pkg/handlers"
)

const portNumber = "8080"
const ipAddress = "127.0.0.1"

var url = fmt.Sprintf("%s:%s", ipAddress, portNumber)

func main() {
	// routing pages
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	// listen to the port for incoming http requests.
	fmt.Printf("Listening traffic at %s\n", url)
	err := http.ListenAndServe(url, nil) // 2nd parameter is `nil` since we didn't send any information to the page.
	log.Fatal(err)
}
