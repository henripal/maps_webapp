package main

import (
	"io"
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

func data(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "WHAT IS THE TEMPLATE????")
}

func main() {
	http.HandleFunc("/data", data)
	http.ListenAndServe(":"+backendPortEnvVariable(), cors.Default().Handler(http.DefaultServeMux))
}
