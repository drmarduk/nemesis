package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	port := flag.Int("port", 8080, "the port we want to listen to")
	host := flag.String("host", "localhost", "the host we want to listen to")
	flag.Parse()

	router := httprouter.New()

	url := fmt.Sprintf("%s:%d", *host, *port)
	log.Fatal(http.ListenAndServe(url, router))
}
