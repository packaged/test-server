package main

import (
	"log"
	"net/http"

	cli "gopkg.in/alecthomas/kingpin.v2"
)

var (
	port      = cli.Flag("port", "port").Short('p').String()
	servePath = cli.Flag("servePath", "serve path").Short('s').Default("./static").String()
	baseURL   = cli.Flag("baseUrl", "base url").Short('b').Default("/").String()
)

func main() {
	cli.Parse()

	if *port == "" {
		log.Fatal("you must specify a port")
	}

	http.Handle(*baseURL, http.StripPrefix(*baseURL, http.FileServer(http.Dir(*servePath))))

	log.Printf("Listening on http://0.0.0.0:" + *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
