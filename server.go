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
	certFile  = cli.Flag("cert", "certificate").Short('c').Default("").String()
	certKey   = cli.Flag("key", "certificate key").Short('k').Default("").String()
)

func main() {
	cli.Parse()

	if *port == "" {
		log.Fatal("you must specify a port")
	}

	http.Handle(*baseURL, http.StripPrefix(*baseURL, http.FileServer(http.Dir(*servePath))))

	log.Printf("Listening on http://0.0.0.0:" + *port)
	if *certFile != "" && *certKey != "" {
		log.Fatal(http.ListenAndServeTLS(":"+*port, *certFile, *certKey, nil))
	} else {
		log.Fatal(http.ListenAndServe(":"+*port, nil))
	}
}
