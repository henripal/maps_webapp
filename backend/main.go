package main

import (
	"net/http"
	"os"

	"github.com/rs/cors"
)

// backendPortEnvVariable returns the port on which the backend
// will run. By default (dev) ":8888"
func backendPortEnvVariable() string {
	env := os.Getenv("BACKENDPORT")
	if env == "" {
		return "8888"
	}
	return env
}

func main() {

	setAPIKey()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080",
			"http://142.93.193.166:80",
			"http://142.93.193.166",
			"http://yourcityfrom.space",
			"http://yourcityfrom.space:80",
			"http://localhost:80",
			"http://localhost"},
		AllowCredentials: true,
	})

	http.HandleFunc("/maps", search)

	http.ListenAndServe(":"+backendPortEnvVariable(), c.Handler(http.DefaultServeMux))
}
