package main

import (
	"net/http"
	"os"

	"webapp_template/backend/auth"
	"webapp_template/backend/usersapi"

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
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	})

	http.HandleFunc("/signup", auth.Signup)
	http.HandleFunc("/signin", auth.Signin)
	http.HandleFunc("/user", usersapi.User)
	http.ListenAndServe(":"+backendPortEnvVariable(), c.Handler(http.DefaultServeMux))
}
