package main

import (
	"flag"

	"github.com/romsnl/httpepe.back/api"
)

var (
	laddr = flag.String("listenaddr", ":3000", "Server listen address")
)

func main() {
	flag.Parse()

	api := api.NewServer()
	api.Serve(*laddr)
}
