package main

import (
	"Ethernaut/contracts/attack"
	"Ethernaut/pkg/client"
	"log"
	"net/url"
)

var localClientURL = url.URL{
	Scheme: "http",
	Host:   "127.0.0.1:8545",
}

func main() {
	c, err := client.New(localClientURL.String())
	if err != nil {
		panic(err)
	}

	if err := attack.Fallback(c); err != nil {
		log.Fatalln(err)
	}
}
