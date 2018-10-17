package main

import (
	"net/http"
	"os"

	"webapp_template/backend/auth"

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
	http.HandleFunc("/signup", auth.Signup)
	http.ListenAndServe(":"+backendPortEnvVariable(), cors.Default().Handler(http.DefaultServeMux))
}
