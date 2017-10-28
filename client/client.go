package client

import (
	"log"
	"net/http"
)

func Healthcheck(port string) {
	url := "http://localhost:" + port
	res, err := http.Get(url + "/health")
	if err != nil {
		log.Fatalf("%s - error %s", url, err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatalf("%s - %s", url, http.StatusText(res.StatusCode))
	}
	log.Printf("%s - OK", url)
}
